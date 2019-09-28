// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrderCreateResponse order create response
// swagger:model OrderCreateResponse
type OrderCreateResponse struct {

	// coupon
	Coupon string `json:"coupon,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// payment url
	// Required: true
	PaymentURL *string `json:"payment_url"`

	// positions
	// Required: true
	Positions []*OrderItem `json:"positions"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this order create response
func (m *OrderCreateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePositions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderCreateResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *OrderCreateResponse) validatePaymentURL(formats strfmt.Registry) error {

	if err := validate.Required("payment_url", "body", m.PaymentURL); err != nil {
		return err
	}

	return nil
}

func (m *OrderCreateResponse) validatePositions(formats strfmt.Registry) error {

	if err := validate.Required("positions", "body", m.Positions); err != nil {
		return err
	}

	for i := 0; i < len(m.Positions); i++ {
		if swag.IsZero(m.Positions[i]) { // not required
			continue
		}

		if m.Positions[i] != nil {
			if err := m.Positions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("positions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrderCreateResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderCreateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderCreateResponse) UnmarshalBinary(b []byte) error {
	var res OrderCreateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
