// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RootResponse root response
//
// swagger:model RootResponse
type RootResponse struct {

	// assets
	Assets string `json:"assets,omitempty"`

	// blocks
	Blocks string `json:"blocks,omitempty"`

	// docs
	Docs string `json:"docs,omitempty"`

	// metadata
	Metadata string `json:"metadata,omitempty"`

	// outputs
	Outputs string `json:"outputs,omitempty"`

	// streams
	Streams string `json:"streams,omitempty"`

	// transactions
	Transactions string `json:"transactions,omitempty"`

	// validators
	Validators string `json:"validators,omitempty"`
}

// Validate validates this root response
func (m *RootResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this root response based on context it is used
func (m *RootResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RootResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RootResponse) UnmarshalBinary(b []byte) error {
	var res RootResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
