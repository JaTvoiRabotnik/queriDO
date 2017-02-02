package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"google.golang.org/appengine"
)

func main() {

	router := NewRouter()

	// Respond to App Engine and Compute Engine health checks.
	// Indicate the server is healthy.
	router.Methods("GET").Path("/_ah/health").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

	// [START request_logging]
	// Delegate all of the HTTP routing and serving to the gorilla/mux router.
	// Log all requests using the standard Apache format.
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
	// [END request_logging]
	appengine.Main()
}
