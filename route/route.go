package route

import (
    "io"
    "net/http"
    "github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func Routes() *mux.Router{
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", hello)
    
    return r
}