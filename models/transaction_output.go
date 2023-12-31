// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TransactionOutput transaction output
//
// swagger:model TransactionOutput
type TransactionOutput struct {

	// amount
	// Required: true
	Amount string `json:"amount"`

	// condition
	// Required: true
	Condition TransactionOutputCondition `json:"condition"`

	// public keys
	// Required: true
	// Min Items: 1
	// Unique: true
	PublicKeys []string `json:"public_keys"`
}

// Validate validates this transaction output
func (m *TransactionOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionOutput) validateAmount(formats strfmt.Registry) error {

	if err := validate.RequiredString("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *TransactionOutput) validateCondition(formats strfmt.Registry) error {

	if err := m.Condition.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("condition")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("condition")
		}
		return err
	}

	return nil
}

func (m *TransactionOutput) validatePublicKeys(formats strfmt.Registry) error {

	if err := validate.Required("public_keys", "body", m.PublicKeys); err != nil {
		return err
	}

	iPublicKeysSize := int64(len(m.PublicKeys))

	if err := validate.MinItems("public_keys", "body", iPublicKeysSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("public_keys", "body", m.PublicKeys); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this transaction output based on the context it is used
func (m *TransactionOutput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCondition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionOutput) contextValidateCondition(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Condition.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("condition")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("condition")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionOutput) UnmarshalBinary(b []byte) error {
	var res TransactionOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TransactionOutputCondition transaction output condition
//
// swagger:model TransactionOutputCondition
type TransactionOutputCondition struct {

	// details
	// Required: true
	Details TransactionOutputConditionDetails `json:"details"`

	// uri
	// Required: true
	// Min Length: 1
	// Pattern: ^ni:///sha-256;([a-zA-Z0-9_-]{0,86})[?](fpt=(ed25519|threshold)-sha-256(&)?|cost=[0-9]+(&)?|subtypes=ed25519-sha-256(&)?){2,3}$
	URI string `json:"uri"`
}

// Validate validates this transaction output condition
func (m *TransactionOutputCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionOutputCondition) validateDetails(formats strfmt.Registry) error {

	if err := m.Details.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("condition" + "." + "details")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("condition" + "." + "details")
		}
		return err
	}

	return nil
}

func (m *TransactionOutputCondition) validateURI(formats strfmt.Registry) error {

	if err := validate.RequiredString("condition"+"."+"uri", "body", m.URI); err != nil {
		return err
	}

	if err := validate.MinLength("condition"+"."+"uri", "body", m.URI, 1); err != nil {
		return err
	}

	if err := validate.Pattern("condition"+"."+"uri", "body", m.URI, `^ni:///sha-256;([a-zA-Z0-9_-]{0,86})[?](fpt=(ed25519|threshold)-sha-256(&)?|cost=[0-9]+(&)?|subtypes=ed25519-sha-256(&)?){2,3}$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this transaction output condition based on the context it is used
func (m *TransactionOutputCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionOutputCondition) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Details.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("condition" + "." + "details")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("condition" + "." + "details")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionOutputCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionOutputCondition) UnmarshalBinary(b []byte) error {
	var res TransactionOutputCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TransactionOutputConditionDetails transaction output condition details
//
// swagger:model TransactionOutputConditionDetails
type TransactionOutputConditionDetails struct {

	// public key
	// Required: true
	// Min Length: 1
	PublicKey string `json:"public_key"`

	// type
	// Required: true
	// Enum: [ed25519-sha-256]
	Type string `json:"type"`
}

// Validate validates this transaction output condition details
func (m *TransactionOutputConditionDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionOutputConditionDetails) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.RequiredString("condition"+"."+"details"+"."+"public_key", "body", m.PublicKey); err != nil {
		return err
	}

	if err := validate.MinLength("condition"+"."+"details"+"."+"public_key", "body", m.PublicKey, 1); err != nil {
		return err
	}

	return nil
}

var transactionOutputConditionDetailsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ed25519-sha-256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionOutputConditionDetailsTypeTypePropEnum = append(transactionOutputConditionDetailsTypeTypePropEnum, v)
	}
}

const (

	// TransactionOutputConditionDetailsTypeEd25519DashShaDash256 captures enum value "ed25519-sha-256"
	TransactionOutputConditionDetailsTypeEd25519DashShaDash256 string = "ed25519-sha-256"
)

// prop value enum
func (m *TransactionOutputConditionDetails) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transactionOutputConditionDetailsTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TransactionOutputConditionDetails) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("condition"+"."+"details"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("condition"+"."+"details"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this transaction output condition details based on context it is used
func (m *TransactionOutputConditionDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TransactionOutputConditionDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionOutputConditionDetails) UnmarshalBinary(b []byte) error {
	var res TransactionOutputConditionDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
