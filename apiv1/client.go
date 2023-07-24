package apiv1

import (
	"context"
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"github.com/filinvadim/bigchaindb-go/client"
	"github.com/filinvadim/bigchaindb-go/client/operations"
	"github.com/filinvadim/bigchaindb-go/models"
	"github.com/filinvadim/stubborn"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/gorilla/websocket"
	"github.com/mr-tron/base58"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type (
	txMode string
	txOp   string

	RESTClientV1 struct {
		ctx    context.Context
		client *client.BigchainClient

		pair *KeyPair
	}

	bcError string

	Transaction      = models.Transaction
	Block            = models.Block
	Fulfills         = models.Fulfills
	Asset            = models.Asset
	Metadata         = models.Metadata
	ValidTransaction = models.ValidTransactionResponse
)

const (
	basePath = "/api/v1"

	ErrNilTx      = bcError("nil transaction provided")
	ErrNoKeyPairs = bcError("no key pairs provided")

	ModeAsync  = txMode(models.TransactionModeAsync)
	ModeSync   = txMode(models.TransactionModeSync)
	ModeCommit = txMode(models.TransactionModeCommit)

	OpCreate   = txOp(models.TransactionOperationCREATE)
	OpTransfer = txOp(models.TransactionOperationTRANSFER)
)

func (e bcError) Error() string {
	return string(e)
}

var DefaultSchemes = []string{"http"}

func NewRESTClientV1(ctx context.Context, host string, ownerKeyPair *KeyPair) (*RESTClientV1, error) {
	if ownerKeyPair == nil {
		return nil, ErrNoKeyPairs
	}
	strings.TrimRight(host, "/")
	_, err := ping(host)
	if err != nil {
		return nil, errors.Wrap(err, "ping")
	}

	transport := httptransport.New(host, basePath, DefaultSchemes)
	cli := client.New(transport, nil)
	return &RESTClientV1{ctx, cli, ownerKeyPair}, err
}

var defaultOpt = func(op *runtime.ClientOperation) {}

type CreateParams struct {
	Mode             txMode
	Amount           float64
	TxData, MetaData any
}

// CreateTx : all key names (e.g. anywhere in the JSON documents stored in asset.data or metadata):
// - must not begin with $
// - must not contain .
// - must not contain the null character (Unicode code point U+0000)
func (c *RESTClientV1) CreateTx(params CreateParams) (*Transaction, error) {
	tx := c.prepareCreateTx(
		strconv.FormatFloat(params.Amount, 'g', -1, 32), params.TxData, params.MetaData,
	)

	err := SignTx(tx, []*KeyPair{c.pair})
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Operations.PostTransactions(&operations.PostTransactionsParams{
		Mode:     func(m txMode) *string { s := string(m); return &s }(params.Mode),
		Context:  c.ctx,
		PostBody: tx,
	}, defaultOpt)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, responseError(resp.Error(), "creation")
	}

	return resp.Payload, nil
}

func (c *RESTClientV1) prepareCreateTx(amount string, td, md any) *Transaction {
	if amount == "0" {
		amount = "1"
	}
	tp := models.TransactionOutputConditionDetailsTypeEd25519DashShaDash256
	tx := models.Transaction{
		Asset: &models.Asset{
			Data: td,
		},

		Metadata:  md,
		Operation: string(OpCreate),
		Outputs: []*models.TransactionOutput{
			{
				Amount: amount,
				Condition: models.TransactionOutputCondition{
					Details: models.TransactionOutputConditionDetails{
						PublicKey: base58.Encode(c.pair.PublicKey),
						Type:      tp,
					},
					URI: NewURIfromKey(c.pair.PublicKey),
				},
				PublicKeys: []string{base58.Encode(c.pair.PublicKey)},
			},
		},
		Version: models.TransactionVersionNr2Dot0,
	}

	tx.Inputs = append(tx.Inputs,
		&models.TransactionInput{
			Fulfillment:  nil,
			Fulfills:     nil,
			OwnersBefore: []string{base58.Encode(c.pair.PublicKey)},
		},
	)

	return &tx
}

type TransferParams struct {
	Mode     txMode
	Amount   float64
	From     Owner
	To       ed25519.PublicKey
	MetaData any
}

type Owner struct {
	TxID string
	Pair *KeyPair
}

func (c *RESTClientV1) TransferOneToOne(params TransferParams) (*Transaction, error) {
	if params.From.TxID == "" {
		return nil, ErrNilTx
	}
	if params.To == nil {
		return nil, errors.New("empty destination public key")
	}
	tx := c.prepareTransferTx(
		strconv.FormatFloat(params.Amount, 'g', -1, 32), params.From, params.To, params.MetaData,
	)

	err := SignTx(tx, []*KeyPair{params.From.Pair})
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Operations.PostTransactions(&operations.PostTransactionsParams{
		Mode:     func(m txMode) *string { s := string(m); return &s }(params.Mode),
		Context:  c.ctx,
		PostBody: tx,
	}, defaultOpt)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, responseError(resp.Error(), "transfer")
	}

	return resp.Payload, nil
}

func (c *RESTClientV1) prepareTransferTx(
	amount string,
	owner Owner,
	dest ed25519.PublicKey,
	md any,
) *Transaction {
	if amount == "0" {
		amount = "1"
	}

	tp := models.TransactionOutputConditionDetailsTypeEd25519DashShaDash256
	tx := models.Transaction{
		Asset: &models.Asset{
			ID: owner.TxID,
		},

		Metadata:  md,
		Operation: string(OpTransfer),
		Outputs: []*models.TransactionOutput{
			{
				Amount: amount,
				Condition: models.TransactionOutputCondition{
					Details: models.TransactionOutputConditionDetails{
						PublicKey: base58.Encode(dest),
						Type:      tp,
					},
					URI: NewURIfromKey(dest),
				},
				PublicKeys: []string{base58.Encode(dest)},
			},
		},
		Version: models.TransactionVersionNr2Dot0,
	}

	tx.Inputs = append(tx.Inputs,
		&models.TransactionInput{
			Fulfillment: nil,
			Fulfills: &models.Fulfills{
				OutputIndex:   0,
				TransactionID: owner.TxID,
			},
			OwnersBefore: []string{base58.Encode(owner.Pair.PublicKey)},
		},
	)

	return &tx
}

// TransferManyToMany TODO
func (c *RESTClientV1) TransferManyToMany(
	mode txMode,
	amount float64,
	prevTx []*Transaction,
	destPubKeys []string,
	metaData any,
) error {
	panic("not implemented")
	return nil
}

func (c *RESTClientV1) GetBlockHeight(blockHeight int64) (*Block, error) {
	resp, err := c.client.Operations.GetBlocksBlockHeight(&operations.GetBlocksBlockHeightParams{
		BlockHeight: blockHeight,
		Context:     c.ctx,
	})
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, responseError(resp.Error(), "block height")
	}

	return resp.Payload, nil
}

func (c *RESTClientV1) GetTransaction(txID string) (*Transaction, error) {
	resp, err := c.client.Operations.GetTransactionsTransactionID(&operations.GetTransactionsTransactionIDParams{
		TransactionID: txID,
		Context:       c.ctx,
	})
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, responseError(resp.Error(), "single tx")
	}

	return resp.Payload, nil
}

func (c *RESTClientV1) ListBlocks(txID string) ([]int64, error) {
	resp, err := c.client.Operations.GetBlocks(&operations.GetBlocksParams{
		TransactionID: txID,
		Context:       c.ctx,
	})
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, responseError(resp.Error(), "blocks list")
	}

	return resp.Payload, nil
}

func (c *RESTClientV1) ListOutputs(pubKey string, spent bool) ([]*Fulfills, error) {
	resp, err := c.client.Operations.GetOutputs(&operations.GetOutputsParams{
		PublicKey: pubKey,
		Spent:     &spent,
		Context:   c.ctx,
	})
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, responseError(resp.Error(), "outputs list")
	}

	return resp.Payload, nil
}

func (c *RESTClientV1) ListTransactions(assetID string, operation txOp, onlyLast bool) ([]*Transaction, error) {
	resp, err := c.client.Operations.GetTransactions(&operations.GetTransactionsParams{
		AssetID:   &assetID,
		LastTx:    &onlyLast,
		Operation: func(op txOp) *string { s := string(op); return &s }(operation),
		Context:   c.ctx,
	})
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, responseError(resp.Error(), "tx list")
	}

	return resp.Payload, nil
}

func (c *RESTClientV1) SearchAsset(search string, limit int64) ([]*Asset, error) {
	var reqLimit *int64
	if limit > 0 {
		reqLimit = &limit
	}
	resp, err := c.client.Operations.GetAssets(&operations.GetAssetsParams{
		Limit:   reqLimit,
		Search:  search,
		Context: c.ctx,
	})
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, responseError(resp.Error(), "search asset")
	}
	return resp.Payload, nil
}

func (c *RESTClientV1) SearchMetadata(search string, limit int64) ([]Metadata, error) {
	var reqLimit *int64
	if limit > 0 {
		reqLimit = &limit
	}
	resp, err := c.client.Operations.GetMetadata(&operations.GetMetadataParams{
		Limit:   reqLimit,
		Search:  search,
		Context: c.ctx,
	})
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, errors.New(resp.Error())
	}
	return resp.Payload, nil
}

func responseError(resp, msg string) error {
	return fmt.Errorf("%s: %s", msg, resp)
}

// ============================================WEBSOCKET=================================================

type WSClient struct {
	url  string
	ctx  context.Context
	logf func(format string, v ...any)

	stub *stubborn.Client
}

type BigChainDBDialer = stubborn.DuplexConnector

func NewWSClient(ctx context.Context, host string, logf func(format string, v ...any)) (*WSClient, error) {
	strings.TrimRight(host, "/")

	return &WSClient{
		url:  "ws://" + host + basePath + "/streams/valid_transactions",
		ctx:  ctx,
		logf: logf,
	}, nil
}

func (ws *WSClient) SubscribeStream(stream chan ValidTransaction) error {
	if ws == nil {
		return errors.New("not connected")
	}

	ws.stub = stubborn.NewStubborn(stubborn.Config{
		IsReconnectable: true,
		MessageType:     stubborn.TextMessage,
		Dialerf: func(ctx context.Context) (BigChainDBDialer, error) {
			conn, _, err := websocket.DefaultDialer.DialContext(
				ctx,
				ws.url,
				nil,
			)
			conn.SetReadDeadline(time.Now().Add(time.Minute))
			conn.WriteMessage(stubborn.TextMessage, []byte("ping"))
			return conn, err
		},
		UnimportantErrs: []error{io.EOF},
	})

	ws.stub.SetErrorHandler(func(err error) {
		if ws.logf != nil {
			ws.logf("websocket connection failure: %v", err)
		}
	})

	ws.stub.SetMessageHandler(func(resp []byte) {
		if ws.logf != nil {
			ws.logf("websocket message: %s", resp)
		}
		var validTx ValidTransaction
		err := json.Unmarshal(resp, &validTx)
		if err != nil {
			if ws.logf != nil {
				ws.logf("websocket message handling failure: %v %s", err, resp)
			}
			return
		}
		stream <- validTx
	})

	return ws.stub.Connect(ws.ctx)
}

func (ws *WSClient) Close() {
	ws.stub.Close()
}

func ping(host string) ([]byte, error) {
	url := "http://" + host + basePath

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bt, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(bt))
	}
	return bt, nil
}
