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

// OrderRequest order request
// swagger:model OrderRequest
type OrderRequest struct {

	// order
	// Required: true
	Order *OrderRequestOrder `json:"order"`

	// user id
	// Required: true
	UserID *string `json:"user_id"`
}

// Validate validates this order request
func (m *OrderRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrder(formats); err != nil {
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

func (m *OrderRequest) validateOrder(formats strfmt.Registry) error {

	if err := validate.Required("order", "body", m.Order); err != nil {
		return err
	}

	if m.Order != nil {
		if err := m.Order.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("order")
			}
			return err
		}
	}

	return nil
}

func (m *OrderRequest) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderRequest) UnmarshalBinary(b []byte) error {
	var res OrderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OrderRequestOrder order request order
// swagger:model OrderRequestOrder
type OrderRequestOrder struct {

	// cafe id
	// Required: true
	CafeID *string `json:"cafe_id"`

	// coupon
	Coupon string `json:"coupon,omitempty"`

	// positions
	// Required: true
	Positions []*OrderItem `json:"positions"`
}

// Validate validates this order request order
func (m *OrderRequestOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCafeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePositions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderRequestOrder) validateCafeID(formats strfmt.Registry) error {

	if err := validate.Required("order"+"."+"cafe_id", "body", m.CafeID); err != nil {
		return err
	}

	return nil
}

func (m *OrderRequestOrder) validatePositions(formats strfmt.Registry) error {

	if err := validate.Required("order"+"."+"positions", "body", m.Positions); err != nil {
		return err
	}

	for i := 0; i < len(m.Positions); i++ {
		if swag.IsZero(m.Positions[i]) { // not required
			continue
		}

		if m.Positions[i] != nil {
			if err := m.Positions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("order" + "." + "positions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderRequestOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderRequestOrder) UnmarshalBinary(b []byte) error {
	var res OrderRequestOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
