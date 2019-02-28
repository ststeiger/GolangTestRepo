
package main

import (
	"fmt"
    "flag"
    //"github.com/gorilla/mux"
    "io"
    "log"
    
    "net"
    "net/http"
    "net/http/fcgi"
    "runtime"
)

var local = flag.String("local", "", "serve as webserver, example: 0.0.0.0:8000")

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())
}

func homeView(w http.ResponseWriter, r *http.Request) {
    headers := w.Header()
    headers.Add("Content-Type", "text/html")
    io.WriteString(w, `<html>
<head>
</head>
<body>
<p>It iis working!</p>
</body>
</html>`)
}




/*
func noDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}
*/

// http://www.dav-muz.net/blog/2013/09/how-to-use-go-and-fastcgi/

// fcgitest -local=":8000"
// http://localhost:8000/
func main() {
    //r := mux.NewRouter()
    //r.HandleFunc("/", homeView)
    
    flag.Parse()
    var err error
    
    if *local != "" { // Run as a local web server
        fmt.Println("Attempting to start web-server")
        //err = http.ListenAndServe(*local, r)
        err = http.ListenAndServe(*local, http.HandlerFunc(homeView))
    } else { // Run as FCGI via standard I/O
		fmt.Println("Attempting to start fastcgi-server")
		
		// Roles:
		// http://www.fastcgi.com/devkit/doc/fastcgi-prog-guide/ch1intro.htm
		
		// http://mwholt.blogspot.ch/2013/05/writing-go-golang-web-app-with-nginx.html
        //err = fcgi.Serve(nil, r)
        //err = fcgi.Serve(nil, http.HandlerFunc(homeView))
        listener, _ := net.Listen("tcp", "127.0.0.1:9001")
        err = fcgi.Serve(listener, http.HandlerFunc(homeView))
    }
    if err != nil {
        log.Fatal(err)
    }
}
