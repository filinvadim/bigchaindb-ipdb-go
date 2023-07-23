package apiv1

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"github.com/filinvadim/bigchaindb-go/models"
	"github.com/kalaspuffar/base64url"
	"golang.org/x/crypto/ed25519"
	"strconv"
	"strings"

	"github.com/go-interledger/cryptoconditions"
	"golang.org/x/crypto/sha3"
)

func newURIfromKey(key ed25519.PublicKey) string {
	res, err := cryptoconditions.NewEd25519Sha256(key, nil)
	if err != nil {
		panic(err)
	}
	return res.Condition().URI()
}

type KeyPair struct {
	PrivateKey ed25519.PrivateKey `json:"privateKey"`
	PublicKey  ed25519.PublicKey  `json:"publicKey"`
}

func NewKeyPair() (*KeyPair, error) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		PublicKey:  pubKey,
		PrivateKey: privKey,
	}, nil
}

const cost = 131072

func NewEd25519Condition(pubKey ed25519.PublicKey) *cryptoconditions.Condition {
	return cryptoconditions.NewSimpleCondition(cryptoconditions.CTEd25519Sha256, pubKey, cost)
}

func SignTx(t *models.Transaction, keyPairs []*KeyPair) error {
	if t == nil {
		return ErrNilTx
	}
	if len(keyPairs) == 0 {
		return ErrNoKeyPairs
	}
	// Set transaction ID to ctnull value
	t.ID = nil

	signedTx := *t

	// Compute signatures of inputs
	for idx, input := range signedTx.Inputs {
		var serializedTx strings.Builder
		s, err := txToString(t)
		if err != nil {
			return err
		}
		serializedTx.WriteString(s)

		keyPair := keyPairs[idx]

		// If fulfills is not empty add to make unique serialization Txn
		if input.Fulfills != nil {
			serializedTx.WriteString(input.Fulfills.TransactionID)
			serializedTx.WriteString(strconv.FormatInt(input.Fulfills.OutputIndex, 10))
		}

		bt := []byte(serializedTx.String())
		h := sha3.New256()
		h.Write(bt)
		hashed := h.Sum(nil)

		// rand reader is ignored within Sign method; crypto.Hash(0) is sanity check to
		// make sure bytes_to_sign is not hashed already - ed25519.PrivateKey cannot sign hashed msg
		signature, err := keyPair.PrivateKey.Sign(rand.Reader, hashed, crypto.Hash(0))

		// https://tools.ietf.org/html/draft-thomas-crypto-conditions-03#section-8.5
		ed25519Fulfillment, err := cryptoconditions.NewEd25519Sha256(keyPair.PublicKey, signature)
		if err != nil {
			return err
		}

		ff, err := ed25519Fulfillment.Encode()
		if err != nil {
			return err
		}
		ffSt := base64url.Encode(ff)
		signedTx.Inputs[idx].Fulfillment = &ffSt
	}

	id, err := createID(signedTx)
	if err != nil {
		return err
	}

	t.Inputs = signedTx.Inputs
	t.ID = &id

	return nil
}

// Create ID of transaction (hash of body)
func createID(t models.Transaction) (string, error) {
	t.ID = nil
	bt, err := txToJSON(t)
	if err != nil {
		return "", err
	}

	// Return hash of serialized txn object
	h := sha3.Sum256(bt)
	// getting slice from array and encode to hex
	return hex.EncodeToString(h[:]), nil
}

func txToString(t *models.Transaction) (string, error) {
	bt, err := txToJSON(*t)
	if err != nil {
		return "", err
	}
	return string(bt), nil
}

// Serialize transaction - encoding/json follows RFC7159 and BDB marshalling
func txToJSON(t models.Transaction) ([]byte, error) {
	bufEncoded := &bytes.Buffer{}
	encoder := json.NewEncoder(bufEncoded)
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(t)
	if err != nil {
		return nil, err
	}
	bt := bufEncoded.Bytes()

	bufCompacted := &bytes.Buffer{}
	err = json.Compact(bufCompacted, bt)
	if err != nil {
		return nil, err
	}
	bt = bufCompacted.Bytes()

	return bt, nil
}
