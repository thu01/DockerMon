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

var templates = template.Must(template.ParseFiles("templates/index.html", 
                                                  "templates/register.html",
                                                  "templates/about.html"))
                                                
type UserInf struct {
    Username string
    Password string
    Email string
}

func Index(w http.ResponseWriter, r *http.Request) {
    cwd, _ := os.Getwd()
    fmt.Println( filepath.Join( cwd, "template/index.html" ) )
    err := templates.ExecuteTemplate(w, "index", nil)
     if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func RegisterPOST(w http.ResponseWriter, r *http.Request) {
    username := r.FormValue("username")
    email := r.FormValue("email")
    password := r.FormValue("password")
    user := &UserInf { username, password, email}
    err := templates.ExecuteTemplate(w, "about", user)
     if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func Routes() *mux.Router{
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    //r.HandleFunc("/", Index)
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./client")))
    r.HandleFunc("/register", RegisterPOST).Methods("POST")
    
    return r
}