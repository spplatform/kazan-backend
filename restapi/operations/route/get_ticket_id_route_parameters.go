// Code generated by go-swagger; DO NOT EDIT.

package route

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTicketIDRouteParams creates a new GetTicketIDRouteParams object
// no default values defined in spec.
func NewGetTicketIDRouteParams() GetTicketIDRouteParams {

	return GetTicketIDRouteParams{}
}

// GetTicketIDRouteParams contains all the bound params for the get ticket ID route operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetTicketIDRoute
type GetTicketIDRouteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ticket ID
	  Required: true
	  Minimum: 1
	  In: path
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTicketIDRouteParams() beforehand.
func (o *GetTicketIDRouteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GetTicketIDRouteParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	if err := o.validateID(formats); err != nil {
		return err
	}

	return nil
}

// validateID carries on validations for parameter ID
func (o *GetTicketIDRouteParams) validateID(formats strfmt.Registry) error {

	if err := validate.Minimum("id", "path", float64(o.ID), 1, false); err != nil {
		return err
	}

	return nil
}
