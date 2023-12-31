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

// NewGetTransactionsTransactionIDParams creates a new GetTransactionsTransactionIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTransactionsTransactionIDParams() *GetTransactionsTransactionIDParams {
	return &GetTransactionsTransactionIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTransactionsTransactionIDParamsWithTimeout creates a new GetTransactionsTransactionIDParams object
// with the ability to set a timeout on a request.
func NewGetTransactionsTransactionIDParamsWithTimeout(timeout time.Duration) *GetTransactionsTransactionIDParams {
	return &GetTransactionsTransactionIDParams{
		timeout: timeout,
	}
}

// NewGetTransactionsTransactionIDParamsWithContext creates a new GetTransactionsTransactionIDParams object
// with the ability to set a context for a request.
func NewGetTransactionsTransactionIDParamsWithContext(ctx context.Context) *GetTransactionsTransactionIDParams {
	return &GetTransactionsTransactionIDParams{
		Context: ctx,
	}
}

// NewGetTransactionsTransactionIDParamsWithHTTPClient creates a new GetTransactionsTransactionIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTransactionsTransactionIDParamsWithHTTPClient(client *http.Client) *GetTransactionsTransactionIDParams {
	return &GetTransactionsTransactionIDParams{
		HTTPClient: client,
	}
}

/*
GetTransactionsTransactionIDParams contains all the parameters to send to the API endpoint

	for the get transactions transaction ID operation.

	Typically these are written to a http.Request.
*/
type GetTransactionsTransactionIDParams struct {

	/* TransactionID.

	   Transaction ID
	*/
	TransactionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get transactions transaction ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTransactionsTransactionIDParams) WithDefaults() *GetTransactionsTransactionIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get transactions transaction ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTransactionsTransactionIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get transactions transaction ID params
func (o *GetTransactionsTransactionIDParams) WithTimeout(timeout time.Duration) *GetTransactionsTransactionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get transactions transaction ID params
func (o *GetTransactionsTransactionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get transactions transaction ID params
func (o *GetTransactionsTransactionIDParams) WithContext(ctx context.Context) *GetTransactionsTransactionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get transactions transaction ID params
func (o *GetTransactionsTransactionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get transactions transaction ID params
func (o *GetTransactionsTransactionIDParams) WithHTTPClient(client *http.Client) *GetTransactionsTransactionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get transactions transaction ID params
func (o *GetTransactionsTransactionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTransactionID adds the transactionID to the get transactions transaction ID params
func (o *GetTransactionsTransactionIDParams) WithTransactionID(transactionID string) *GetTransactionsTransactionIDParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get transactions transaction ID params
func (o *GetTransactionsTransactionIDParams) SetTransactionID(transactionID string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTransactionsTransactionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param transaction_id
	if err := r.SetPathParam("transaction_id", o.TransactionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
