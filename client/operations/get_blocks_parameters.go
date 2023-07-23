// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetBlocksParams creates a new GetBlocksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBlocksParams() *GetBlocksParams {
	return &GetBlocksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBlocksParamsWithTimeout creates a new GetBlocksParams object
// with the ability to set a timeout on a request.
func NewGetBlocksParamsWithTimeout(timeout time.Duration) *GetBlocksParams {
	return &GetBlocksParams{
		timeout: timeout,
	}
}

// NewGetBlocksParamsWithContext creates a new GetBlocksParams object
// with the ability to set a context for a request.
func NewGetBlocksParamsWithContext(ctx context.Context) *GetBlocksParams {
	return &GetBlocksParams{
		Context: ctx,
	}
}

// NewGetBlocksParamsWithHTTPClient creates a new GetBlocksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBlocksParamsWithHTTPClient(client *http.Client) *GetBlocksParams {
	return &GetBlocksParams{
		HTTPClient: client,
	}
}

/*
GetBlocksParams contains all the parameters to send to the API endpoint

	for the get blocks operation.

	Typically these are written to a http.Request.
*/
type GetBlocksParams struct {

	/* TransactionID.

	   Transaction ID
	*/
	TransactionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get blocks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBlocksParams) WithDefaults() *GetBlocksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get blocks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBlocksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get blocks params
func (o *GetBlocksParams) WithTimeout(timeout time.Duration) *GetBlocksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get blocks params
func (o *GetBlocksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get blocks params
func (o *GetBlocksParams) WithContext(ctx context.Context) *GetBlocksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get blocks params
func (o *GetBlocksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get blocks params
func (o *GetBlocksParams) WithHTTPClient(client *http.Client) *GetBlocksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get blocks params
func (o *GetBlocksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTransactionID adds the transactionID to the get blocks params
func (o *GetBlocksParams) WithTransactionID(transactionID string) *GetBlocksParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get blocks params
func (o *GetBlocksParams) SetTransactionID(transactionID string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBlocksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param transaction_id
	qrTransactionID := o.TransactionID
	qTransactionID := qrTransactionID
	if qTransactionID != "" {

		if err := r.SetQueryParam("transaction_id", qTransactionID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
