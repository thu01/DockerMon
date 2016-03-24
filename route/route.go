package route

import (
    "html/template"
    "net/http"
    //"path"
    "os"
    "path/filepath"
    "fmt"
    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    cwd, _ := os.Getwd()
    fmt.Println( filepath.Join( cwd, "template/index.html" ) )
    templates := template.Must(template.ParseFiles("templates/index.html"))
    err := templates.ExecuteTemplate(w, "index", nil)
     if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func Routes() *mux.Router{
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", Index)
    
    return r
}