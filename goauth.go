
package main


import (
    "fmt"
    "log"
    "net/http"

    "github.com/abbot/go-http-auth"
)


func Secret(user, realm string) string {
    users := map[string]string{
        "john": "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1", //hello
    }

    if a, ok := users[user]; ok {
        return a
    }
    return ""
}


func doRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>static file server</h1><p><a href='./static'>folder</p>")
}


func handleFileServer(w http.ResponseWriter, r *http.Request) {
    fs := http.FileServer(http.Dir("static"))
    http.StripPrefix("/static/", fs)
}


/*
// You need to return StripPrefix's ServeHTTP method, for example:
func handleFileServer(dir, prefix string) http.HandlerFunc {
    fs := http.FileServer(http.Dir(dir))
    realHandler := http.StripPrefix(prefix, fs).ServeHTTP
    return func(w http.ResponseWriter, req *http.Request) {
        log.Println(req.URL)
        realHandler(w, req)
    }
}

func main()
    //....
    http.HandleFunc("/static/", auth.JustCheck(authenticator, handleFileServer("/tmp", "/static/")))
    //....
}
*/


// http://stackoverflow.com/questions/25552107/golang-how-to-serve-static-files-with-basic-authentication

func main() {
    authenticator := auth.NewBasicAuthenticator("localhost", Secret)
	
    // how to secure the FileServer with basic authentication??
    // fs := http.FileServer(http.Dir("static"))
    // http.Handle("/static/", http.StripPrefix("/static/", fs))
    http.HandleFunc("/static/", auth.JustCheck(authenticator, handleFileServer))

    http.HandleFunc("/", auth.JustCheck(authenticator, doRoot))

    log.Println(`Listening... http://localhost:3000
 folder is ./static
 authentication in map users`)
    http.ListenAndServe(":3001", nil)
}



// https://github.com/go-authboss/authboss
// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/14.4.html

// https://mschoebel.info/2014/03/09/snippet-golang-webapp-login-logout/
// https://github.com/golang/oauth2
