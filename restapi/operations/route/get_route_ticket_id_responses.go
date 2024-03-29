// Code generated by go-swagger; DO NOT EDIT.

package route

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/spplatform/kazan-backend/models"
)

// GetRouteTicketIDOKCode is the HTTP code returned for type GetRouteTicketIDOK
const GetRouteTicketIDOKCode int = 200

/*GetRouteTicketIDOK OK

swagger:response getRouteTicketIdOK
*/
type GetRouteTicketIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.RouteResponse `json:"body,omitempty"`
}

// NewGetRouteTicketIDOK creates GetRouteTicketIDOK with default headers values
func NewGetRouteTicketIDOK() *GetRouteTicketIDOK {

	return &GetRouteTicketIDOK{}
}

// WithPayload adds the payload to the get route ticket Id o k response
func (o *GetRouteTicketIDOK) WithPayload(payload *models.RouteResponse) *GetRouteTicketIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get route ticket Id o k response
func (o *GetRouteTicketIDOK) SetPayload(payload *models.RouteResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRouteTicketIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRouteTicketIDBadRequestCode is the HTTP code returned for type GetRouteTicketIDBadRequest
const GetRouteTicketIDBadRequestCode int = 400

/*GetRouteTicketIDBadRequest Bad request

swagger:response getRouteTicketIdBadRequest
*/
type GetRouteTicketIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.StatusResponse `json:"body,omitempty"`
}

// NewGetRouteTicketIDBadRequest creates GetRouteTicketIDBadRequest with default headers values
func NewGetRouteTicketIDBadRequest() *GetRouteTicketIDBadRequest {

	return &GetRouteTicketIDBadRequest{}
}

// WithPayload adds the payload to the get route ticket Id bad request response
func (o *GetRouteTicketIDBadRequest) WithPayload(payload *models.StatusResponse) *GetRouteTicketIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get route ticket Id bad request response
func (o *GetRouteTicketIDBadRequest) SetPayload(payload *models.StatusResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRouteTicketIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRouteTicketIDNotFoundCode is the HTTP code returned for type GetRouteTicketIDNotFound
const GetRouteTicketIDNotFoundCode int = 404

/*GetRouteTicketIDNotFound Not found

swagger:response getRouteTicketIdNotFound
*/
type GetRouteTicketIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.StatusResponse `json:"body,omitempty"`
}

// NewGetRouteTicketIDNotFound creates GetRouteTicketIDNotFound with default headers values
func NewGetRouteTicketIDNotFound() *GetRouteTicketIDNotFound {

	return &GetRouteTicketIDNotFound{}
}

// WithPayload adds the payload to the get route ticket Id not found response
func (o *GetRouteTicketIDNotFound) WithPayload(payload *models.StatusResponse) *GetRouteTicketIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get route ticket Id not found response
func (o *GetRouteTicketIDNotFound) SetPayload(payload *models.StatusResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRouteTicketIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRouteTicketIDInternalServerErrorCode is the HTTP code returned for type GetRouteTicketIDInternalServerError
const GetRouteTicketIDInternalServerErrorCode int = 500

/*GetRouteTicketIDInternalServerError Internal server error

swagger:response getRouteTicketIdInternalServerError
*/
type GetRouteTicketIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.StatusResponse `json:"body,omitempty"`
}

// NewGetRouteTicketIDInternalServerError creates GetRouteTicketIDInternalServerError with default headers values
func NewGetRouteTicketIDInternalServerError() *GetRouteTicketIDInternalServerError {

	return &GetRouteTicketIDInternalServerError{}
}

// WithPayload adds the payload to the get route ticket Id internal server error response
func (o *GetRouteTicketIDInternalServerError) WithPayload(payload *models.StatusResponse) *GetRouteTicketIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get route ticket Id internal server error response
func (o *GetRouteTicketIDInternalServerError) SetPayload(payload *models.StatusResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRouteTicketIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
