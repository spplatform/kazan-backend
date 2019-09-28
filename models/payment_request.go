// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaymentRequest payment request
// swagger:model PaymentRequest
type PaymentRequest struct {

	// payment id
	// Required: true
	PaymentID *string `json:"payment_id"`

	// user id
	// Required: true
	UserID *string `json:"user_id"`
}

// Validate validates this payment request
func (m *PaymentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRequest) validatePaymentID(formats strfmt.Registry) error {

	if err := validate.Required("payment_id", "body", m.PaymentID); err != nil {
		return err
	}

	return nil
}

func (m *PaymentRequest) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRequest) UnmarshalBinary(b []byte) error {
	var res PaymentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}