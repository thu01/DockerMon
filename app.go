package main

import (
	"github.com/thu01/GoWebApp/route"
	"net/http"
)

func init() {
}

func main() {

	// Bind to a port and pass our router in
	http.ListenAndServe(":8000", route.Routes())
}
