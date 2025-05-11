package main

import (
	"fmt"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

// $ curl -i localhost:4000/v1/healthcheck
// HTTP/1.1 200 OK
// Date: Mon, 05 Apr 2021 17:46:14 GMT
// Content-Length: 58
// Content-Type: text/plain; charset=utf-8

// status: available
// environment: development
// version: 1.0.0
