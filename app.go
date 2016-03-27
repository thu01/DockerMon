package main

import (
    "net/http"
	"github.com/thu01/GoWebApp/route"
)


func init() {
}

func main() {

    // Bind to a port and pass our router in
    http.ListenAndServe(":8000", route.Routes())
}
