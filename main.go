package main

import (
	"net/http"

	"github.com/derstruct/doors-tutorial/catalog"
	"github.com/derstruct/doors-tutorial/home"
	"github.com/doors-dev/doors"
)

func main() {
	r := doors.NewRouter()

	r.Use(
		doors.UsePage(home.Handler),
		// our new catalog page
		doors.UsePage(catalog.Handler),
	)

	err := http.ListenAndServeTLS(":8443", "localhost+2.pem", "localhost+2-key.pem", r)
	if err != nil {
		panic(err)
	}
}
