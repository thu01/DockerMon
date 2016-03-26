package route

import (
    "encoding/json"
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

type Response struct {
    Code int
    Message interface{}
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
    fmt.Println("RegisterPOST")
    
    username := r.FormValue("username")
    email := r.FormValue("email")
    password := r.FormValue("password")
    user := &UserInf { username, password, email}
    statusCode := http.StatusOK
    userJson, err := json.Marshal(user)
    if err != nil {
        statusCode = http.StatusInternalServerError
    }
    
    fmt.Println(string(userJson))
    response := &Response{statusCode, userJson}
    responseJson, err := json.Marshal(response)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(responseJson))
    w.Write(responseJson)
}

func Routes() *mux.Router{
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    //r.HandleFunc("/", Index)
    r.HandleFunc("/api/users", RegisterPOST).Methods("POST")
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./client")))
    return r
}