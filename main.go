package main

import (
	"github.com/derstruct/doors-tutorial/home"
	"github.com/doors-dev/doors"
	"net/http"
)

func main() {
	// create doors router
	r := doors.NewRouter()

	// attach home handler
	r.Use(doors.UsePage(home.Handler))

	// start server with our self signed cert
	err := http.ListenAndServeTLS(":8443", "localhost+2.pem", "localhost+2-key.pem", r)
	if err != nil {
		panic(err)
	}
}
