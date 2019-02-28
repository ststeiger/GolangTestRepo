
package main


import (
     "fmt"
     "net/http"
)


func someHandler(w http.ResponseWriter, r *http.Request) {

     if r.URL.String() == "/favicon.ico" {
	return
     }

     fmt.Println("Scheme: " + r.URL.Scheme)
     fmt.Println("Host: " + r.URL.Host)
     fmt.Println("Path: " + r.URL.Path)

     // https://golang.org/pkg/net/url/
     // scheme://[userinfo@]host/path[?query][#fragment]
     fmt.Println("RawPath: " + r.URL.RawPath)  // encoded path hint (Go 1.5 and later only; see EscapedPath method)
     fmt.Println("RawQuery: " + r.URL.RawQuery) // encoded query values, without '?'
     fmt.Println("Fragment: " + r.URL.Fragment) // fragment for references, without '#
     fmt.Println("Opaque: " + r.URL.Opaque) // encoded opaque data

	 

     fmt.Println("Original URL: " + r.URL.String())
     fmt.Println("Request URI: " + r.RequestURI)
}

// http://grokbase.com/t/gg/golang-nuts/144j758twr/go-nuts-is-the-full-absolute-url-of-an-http-request-directly-available

func main() {
     http.HandleFunc("/", someHandler)
     http.ListenAndServe(":8080", nil)
}

