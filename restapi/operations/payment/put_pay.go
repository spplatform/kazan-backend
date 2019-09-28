// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutPayHandlerFunc turns a function with the right signature into a put pay handler
type PutPayHandlerFunc func(PutPayParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutPayHandlerFunc) Handle(params PutPayParams) middleware.Responder {
	return fn(params)
}

// PutPayHandler interface for that can handle valid put pay params
type PutPayHandler interface {
	Handle(PutPayParams) middleware.Responder
}

// NewPutPay creates a new http.Handler for the put pay operation
func NewPutPay(ctx *middleware.Context, handler PutPayHandler) *PutPay {
	return &PutPay{Context: ctx, Handler: handler}
}

/*PutPay swagger:route PUT /pay payment putPay

get route by ticket number

*/
type PutPay struct {
	Context *middleware.Context
	Handler PutPayHandler
}

func (o *PutPay) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutPayParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}