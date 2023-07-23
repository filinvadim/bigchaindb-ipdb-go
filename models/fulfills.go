// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Fulfills fulfills
//
// swagger:model Fulfills
type Fulfills struct {

	// output index
	// Required: true
	OutputIndex int64 `json:"output_index"`

	// transaction id
	// Required: true
	TransactionID string `json:"transaction_id"`
}

// Validate validates this fulfills
func (m *Fulfills) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutputIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Fulfills) validateOutputIndex(formats strfmt.Registry) error {

	if err := validate.Required("output_index", "body", int64(m.OutputIndex)); err != nil {
		return err
	}

	return nil
}

func (m *Fulfills) validateTransactionID(formats strfmt.Registry) error {

	if err := validate.RequiredString("transaction_id", "body", m.TransactionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this fulfills based on context it is used
func (m *Fulfills) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Fulfills) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Fulfills) UnmarshalBinary(b []byte) error {
	var res Fulfills
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}