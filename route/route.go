package route

import (
    "encoding/json"
    "net/http"
    //"path"
    "fmt"
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
)

const (
	   USER_STATUS_NOT_EXISTS = iota
       USER_STATUS_NORMAL     
)

type UserInf struct {
    Status int
    Username string
    Password string
    Email string
}

type Response struct {
    Status int
    Message interface{}
}

var sessionStore = sessions.NewCookieStore([]byte("GoWebApp-Session-Store"))

func init() {
    sessionStore.Options = &sessions.Options{
        Domain:   "localhost",
        Path:     "/",
        MaxAge:   3600 * 8, // 8 hours
        HttpOnly: true,
    }
}

func NotFound(w http.ResponseWriter, r *http.Request) {
    fmt.Println("NotFound")
    //TODO: only serve index.html for exposed url, otherwise serve error.html
    http.ServeFile(w, r, "./client/index.html")
}

func RegisterPOST(w http.ResponseWriter, r *http.Request) {
    fmt.Println("RegisterPOST")
    
    username := r.FormValue("username")
    email := r.FormValue("email")
    password := r.FormValue("password")
    user := &UserInf { USER_STATUS_NORMAL, username, password, email}
    statusCode := http.StatusOK
    
    response := &Response{statusCode, user}
    
    responseJson, err := json.Marshal(response)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(responseJson))
    w.Write(responseJson)
}

func UserGET(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    username := vars["username"]
    userMap := map[string]UserInf{
        "thu": UserInf{USER_STATUS_NORMAL, "thu", "12345", "thu@gmail.com"},
        "lss": UserInf{USER_STATUS_NORMAL, "lss", "12345", "lss@gmail.com"},
    }
    
    user, ok := userMap[username]
    if !ok {
        user = UserInf{Status:USER_STATUS_NOT_EXISTS}
    }
    statusCode := http.StatusOK
    response := &Response{statusCode, user}
    responseJson, err := json.Marshal(response)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(responseJson))
    w.Write(responseJson)
}

func SessionPOST(w http.ResponseWriter, r *http.Request) {
    userMap := map[string]UserInf{
        "thu": UserInf{USER_STATUS_NORMAL, "thu", "12345", "thu@gmail.com"},
        "lss": UserInf{USER_STATUS_NORMAL, "lss", "12345", "lss@gmail.com"},
    }

    statusCode := http.StatusOK
    session, err := sessionStore.Get(r, "GoWebApp-Login-Session");
    if err!=nil {
        fmt.Println(err)
        response := &Response{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)}
        responseJson, _ := json.Marshal(response)
        w.Write(responseJson)
        return
    }
    
    username := r.FormValue("username")
    password := r.FormValue("password")
    for k,v := range session.Values {
        fmt.Println(k, "=", v)
    } 
    if session.Values["username"]==username {
        user, _ := userMap[username]
        response := &Response{statusCode, user}
        responseJson, err := json.Marshal(response)
        if err != nil {
        	    fmt.Println(err)
        }
        w.Write(responseJson)
        return
    }
    
    user, ok := userMap[username]
    if !ok {
        statusCode = http.StatusNotFound
        response := &Response{statusCode, http.StatusText(statusCode)}
        responseJson, _ := json.Marshal(response)
        w.Write(responseJson)
        return
    }
    if user.Password != password {
        statusCode = http.StatusUnauthorized
        response := &Response{statusCode, http.StatusText(statusCode)}
        responseJson, _ := json.Marshal(response)
        w.Write(responseJson)
        return
        
    }
    response := &Response{statusCode, user}
    responseJson, err := json.Marshal(response)
    if err != nil {
        fmt.Println(err)
    }
    session.Values["username"]=username
    session.Save(r,w)
    w.Write(responseJson)
}

func Routes() *mux.Router{
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/api/users", RegisterPOST).Methods("POST")
    //TODO: Add more strict checking for username
    r.HandleFunc("/api/users/{username:[0-9a-zA-z._]+}", UserGET).Methods("GET")
    r.HandleFunc("/api/session", SessionPOST).Methods("POST")
    
    //Route for static files
    r.Handle("/", http.FileServer(http.Dir("./client/")))
    r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", 
                                 http.FileServer(http.Dir("./client/js/"))))
    r.PathPrefix("/html/").Handler(http.StripPrefix("/html/",
                                   http.FileServer(http.Dir("./client/html/"))))
    r.PathPrefix("/bower_components/").Handler(http.StripPrefix("/bower_components/",
                                               http.FileServer(http.Dir("./client/bower_components/"))))
                                               
    //Route for not-found url
    r.NotFoundHandler = http.HandlerFunc(NotFound)
    
    return r
}