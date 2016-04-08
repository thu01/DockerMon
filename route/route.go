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
    Status int	 `json:"status"` 
    Username string `json:"username"`
    Password string	`json:"-"` 
    Email string    `json:"email"`
}

type Response struct {
    Status int
    Message interface{}
}

type HttpStatus struct {
    Code int
    Message string
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

func WriteResponse(w http.ResponseWriter, status int, message interface{}) {
    resp := &Response{status, message}
    respJson, err := json.Marshal(resp)
    if err!=nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
    }
    w.WriteHeader(status)
    w.Write(respJson)
}

func GetUserInfo(username string) (UserInf, HttpStatus) {
    code := http.StatusOK
    status := HttpStatus{code, http.StatusText(code)}
    
    //TODO: Communicate with db
    userMap := map[string]UserInf{
        "thu": UserInf{USER_STATUS_NORMAL, "thu", "12345", "thu@gmail.com"},
        "lss": UserInf{USER_STATUS_NORMAL, "lss", "12345", "lss@gmail.com"},
    }
    user, ok := userMap[username]
    if !ok {
        code = http.StatusNotFound
        return UserInf{}, HttpStatus{code, http.StatusText(code)}
    }

    return user, status
}

func UserExists(username string) (bool) {
    //TODO: Communicate with db
    userMap := map[string]UserInf{
        "thu": UserInf{USER_STATUS_NORMAL, "thu", "12345", "thu@gmail.com"},
        "lss": UserInf{USER_STATUS_NORMAL, "lss", "12345", "lss@gmail.com"},
    }
    _, ok := userMap[username]
    if !ok {
        return false
    } else {
        return true
    }
}

func NotFound(w http.ResponseWriter, r *http.Request) {
    fmt.Println("NotFound")
    //TODO: only serve index.html for exposed url, otherwise serve error.html
    http.ServeFile(w, r, "./client/index.html")
}

func RegisterPOST(w http.ResponseWriter, r *http.Request) {
    //TODO: Add log
    fmt.Println("RegisterPOST")
    code := http.StatusOK
   
    username := r.FormValue("username")
    email := r.FormValue("email")
    password := r.FormValue("password")

    if UserExists(username)==true {
        code = http.StatusConflict
        WriteResponse(w, code, HttpStatus{code, http.StatusText(code)})
        return
    }
    //TODO: Check if email has been used
    //TODO: Save session

    user := &UserInf { USER_STATUS_NORMAL, username, password, email}
    WriteResponse(w, code, user)
}

func UserGET(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    username := vars["username"]

    user, status := GetUserInfo(username)
    if status.Code!=http.StatusOK {
        WriteResponse(w, status.Code, status)
        return
    }
    
    WriteResponse(w, status.Code, user)
}

func SessionPOST(w http.ResponseWriter, r *http.Request) {
    code := http.StatusOK
    status := HttpStatus{code, http.StatusText(code)}

    session, err := sessionStore.Get(r, "GoWebApp-Login-Session");
    if err!=nil {
        code = http.StatusInternalServerError
        WriteResponse(w, code, HttpStatus{code, http.StatusText(code)}) 
        return
    }
    
    username := r.FormValue("username")
    password := r.FormValue("password")
    if session.Values["username"]==username {
        user, status := GetUserInfo(username)
        if status.Code!=http.StatusOK {
            WriteResponse(w, status.Code, status)
        }
        
        WriteResponse(w, status.Code, user)
        return
    }

    user, status := GetUserInfo(username)
    if status.Code!=http.StatusOK {
        WriteResponse(w, status.Code, status)
        return
    }
	if password != user.Password {
        code = http.StatusUnauthorized
        WriteResponse(w, code, HttpStatus{code, http.StatusText(code)}) 
        return
    }
    
    session.Values["username"]=username
    session.Save(r,w)
    WriteResponse(w, status.Code, user)
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