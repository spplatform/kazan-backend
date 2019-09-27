// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/spplatform/kazan-backend/handlers"
	"github.com/spplatform/kazan-backend/restapi/operations"
	"github.com/spplatform/kazan-backend/restapi/operations/order"
	"github.com/spplatform/kazan-backend/restapi/operations/route"
	"log"
	"net/http"
	"os"
)

//go:generate swagger generate server --target ../../kazan --name Kazan --spec ../swagger.yml

var hdlr *handlers.Handler

func configureFlags(api *operations.KazanAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.KazanAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	var err error
	if hdlr == nil {
		log.Print("configureAPI: init handler")
		user := os.Getenv("MGO_USER")
		pwd := os.Getenv("MGO_PASS")
		host := os.Getenv("MGO_HOST")
		database := os.Getenv("MGO_DATABASE")
		hdlr, err = handlers.NewHandler(user, pwd, host, database)
		if err != nil {
			log.Print("can't connect to mongo: ", err)
		} else {
			log.Print("register handlers")
			api.OrderGetOrderIDHandler = order.GetOrderIDHandlerFunc(hdlr.HandleGetOrder)
			api.OrderPostOrderHandler = order.PostOrderHandlerFunc(hdlr.HandlePostOrder)
			api.RouteGetTicketIDRouteHandler = route.GetTicketIDRouteHandlerFunc(hdlr.HandleGetTicketRoute)
		}
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {

	log.Printf("configureServer")
	//instantiate server
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
