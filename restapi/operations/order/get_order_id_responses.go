// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/spplatform/kazan-backend/models"
)

// GetOrderIDOKCode is the HTTP code returned for type GetOrderIDOK
const GetOrderIDOKCode int = 200

/*GetOrderIDOK OK

swagger:response getOrderIdOK
*/
type GetOrderIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.OrderResponse `json:"body,omitempty"`
}

// NewGetOrderIDOK creates GetOrderIDOK with default headers values
func NewGetOrderIDOK() *GetOrderIDOK {

	return &GetOrderIDOK{}
}

// WithPayload adds the payload to the get order Id o k response
func (o *GetOrderIDOK) WithPayload(payload *models.OrderResponse) *GetOrderIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get order Id o k response
func (o *GetOrderIDOK) SetPayload(payload *models.OrderResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrderIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOrderIDBadRequestCode is the HTTP code returned for type GetOrderIDBadRequest
const GetOrderIDBadRequestCode int = 400

/*GetOrderIDBadRequest Bad request

swagger:response getOrderIdBadRequest
*/
type GetOrderIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.StatusResponse `json:"body,omitempty"`
}

// NewGetOrderIDBadRequest creates GetOrderIDBadRequest with default headers values
func NewGetOrderIDBadRequest() *GetOrderIDBadRequest {

	return &GetOrderIDBadRequest{}
}

// WithPayload adds the payload to the get order Id bad request response
func (o *GetOrderIDBadRequest) WithPayload(payload *models.StatusResponse) *GetOrderIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get order Id bad request response
func (o *GetOrderIDBadRequest) SetPayload(payload *models.StatusResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrderIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOrderIDNotFoundCode is the HTTP code returned for type GetOrderIDNotFound
const GetOrderIDNotFoundCode int = 404

/*GetOrderIDNotFound Not found

swagger:response getOrderIdNotFound
*/
type GetOrderIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.StatusResponse `json:"body,omitempty"`
}

// NewGetOrderIDNotFound creates GetOrderIDNotFound with default headers values
func NewGetOrderIDNotFound() *GetOrderIDNotFound {

	return &GetOrderIDNotFound{}
}

// WithPayload adds the payload to the get order Id not found response
func (o *GetOrderIDNotFound) WithPayload(payload *models.StatusResponse) *GetOrderIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get order Id not found response
func (o *GetOrderIDNotFound) SetPayload(payload *models.StatusResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrderIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOrderIDInternalServerErrorCode is the HTTP code returned for type GetOrderIDInternalServerError
const GetOrderIDInternalServerErrorCode int = 500

/*GetOrderIDInternalServerError Internal server error

swagger:response getOrderIdInternalServerError
*/
type GetOrderIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.StatusResponse `json:"body,omitempty"`
}

// NewGetOrderIDInternalServerError creates GetOrderIDInternalServerError with default headers values
func NewGetOrderIDInternalServerError() *GetOrderIDInternalServerError {

	return &GetOrderIDInternalServerError{}
}

// WithPayload adds the payload to the get order Id internal server error response
func (o *GetOrderIDInternalServerError) WithPayload(payload *models.StatusResponse) *GetOrderIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get order Id internal server error response
func (o *GetOrderIDInternalServerError) SetPayload(payload *models.StatusResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrderIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
