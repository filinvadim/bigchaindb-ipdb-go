// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Transaction transaction
//
// swagger:model Transaction
type Transaction struct {

	// asset
	// Required: true
	Asset *Asset `json:"asset"`

	// id
	// Required: true
	// Min Length: 32
	ID *string `json:"id"`

	// inputs
	// Required: true
	// Min Items: 1
	Inputs []*TransactionInput `json:"inputs"`

	// metadata
	// Required: true
	Metadata Metadata `json:"metadata"`

	// operation
	// Required: true
	// Enum: [CREATE TRANSFER VALIDATOR_ELECTION CHAIN_MIGRATION_ELECTION VOTE]
	Operation string `json:"operation"`

	// outputs
	// Required: true
	Outputs []*TransactionOutput `json:"outputs"`

	// version
	// Required: true
	// Enum: [2.0]
	Version string `json:"version"`
}

// Validate validates this transaction
func (m *Transaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transaction) validateAsset(formats strfmt.Registry) error {

	if err := validate.Required("asset", "body", m.Asset); err != nil {
		return err
	}

	if m.Asset != nil {
		if err := m.Asset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("asset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("asset")
			}
			return err
		}
	}

	return nil
}

func (m *Transaction) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", *m.ID, 32); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateInputs(formats strfmt.Registry) error {

	if err := validate.Required("inputs", "body", m.Inputs); err != nil {
		return err
	}

	iInputsSize := int64(len(m.Inputs))

	if err := validate.MinItems("inputs", "body", iInputsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Inputs); i++ {
		if swag.IsZero(m.Inputs[i]) { // not required
			continue
		}

		if m.Inputs[i] != nil {
			if err := m.Inputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Transaction) validateMetadata(formats strfmt.Registry) error {

	if m.Metadata == nil {
		return errors.Required("metadata", "body", nil)
	}

	return nil
}

var transactionTypeOperationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATE","TRANSFER","VALIDATOR_ELECTION","CHAIN_MIGRATION_ELECTION","VOTE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionTypeOperationPropEnum = append(transactionTypeOperationPropEnum, v)
	}
}

const (

	// TransactionOperationCREATE captures enum value "CREATE"
	TransactionOperationCREATE string = "CREATE"

	// TransactionOperationTRANSFER captures enum value "TRANSFER"
	TransactionOperationTRANSFER string = "TRANSFER"

	// TransactionOperationVALIDATORELECTION captures enum value "VALIDATOR_ELECTION"
	TransactionOperationVALIDATORELECTION string = "VALIDATOR_ELECTION"

	// TransactionOperationCHAINMIGRATIONELECTION captures enum value "CHAIN_MIGRATION_ELECTION"
	TransactionOperationCHAINMIGRATIONELECTION string = "CHAIN_MIGRATION_ELECTION"

	// TransactionOperationVOTE captures enum value "VOTE"
	TransactionOperationVOTE string = "VOTE"
)

// prop value enum
func (m *Transaction) validateOperationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transactionTypeOperationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Transaction) validateOperation(formats strfmt.Registry) error {

	if err := validate.RequiredString("operation", "body", m.Operation); err != nil {
		return err
	}

	// value enum
	if err := m.validateOperationEnum("operation", "body", m.Operation); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateOutputs(formats strfmt.Registry) error {

	if err := validate.Required("outputs", "body", m.Outputs); err != nil {
		return err
	}

	for i := 0; i < len(m.Outputs); i++ {
		if swag.IsZero(m.Outputs[i]) { // not required
			continue
		}

		if m.Outputs[i] != nil {
			if err := m.Outputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var transactionTypeVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["2.0"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionTypeVersionPropEnum = append(transactionTypeVersionPropEnum, v)
	}
}

const (

	// TransactionVersionNr2Dot0 captures enum value "2.0"
	TransactionVersionNr2Dot0 string = "2.0"
)

// prop value enum
func (m *Transaction) validateVersionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transactionTypeVersionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Transaction) validateVersion(formats strfmt.Registry) error {

	if err := validate.RequiredString("version", "body", m.Version); err != nil {
		return err
	}

	// value enum
	if err := m.validateVersionEnum("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this transaction based on the context it is used
func (m *Transaction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAsset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transaction) contextValidateAsset(ctx context.Context, formats strfmt.Registry) error {

	if m.Asset != nil {

		if err := m.Asset.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("asset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("asset")
			}
			return err
		}
	}

	return nil
}

func (m *Transaction) contextValidateInputs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Inputs); i++ {

		if m.Inputs[i] != nil {

			if swag.IsZero(m.Inputs[i]) { // not required
				return nil
			}

			if err := m.Inputs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Transaction) contextValidateOutputs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Outputs); i++ {

		if m.Outputs[i] != nil {

			if swag.IsZero(m.Outputs[i]) { // not required
				return nil
			}

			if err := m.Outputs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Transaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Transaction) UnmarshalBinary(b []byte) error {
	var res Transaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
