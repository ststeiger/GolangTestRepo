
package main 

// http://capotej.com/blog/2013/10/07/golang-http-handlers-as-middleware/


import "fmt"
import "net/http"


func OurLoggingHandler(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Println(*r.URL)
    h.ServeHTTP(w, r)
  })
}

func main() {
    fileHandler := http.FileServer(http.Dir("/tmp"))
    wrappedHandler := OurLoggingHandler(fileHandler)
    http.ListenAndServe(":8080", wrappedHandler)
}

