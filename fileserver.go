
package main


import (
    "fmt"
    "net/http"
    "strings"
)   

var chttp = http.NewServeMux()


func main() {
    chttp.Handle("/", http.FileServer(http.Dir("./")))


    // http://stackoverflow.com/questions/17541333/fileserver-handler-with-some-other-http-handlers
    //http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("."))))


    http.HandleFunc("/", HomeHandler) // homepage
    http.ListenAndServe(":8080", nil)
}   


// http://stackoverflow.com/questions/14086063/serve-homepage-and-static-content-from-root
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)

    if (strings.Contains(r.URL.Path, ".")) {
        chttp.ServeHTTP(w, r)
	return
    } 
    




    fmt.Fprintf(w, "HomeHandler")
} 


// http://www.shakedos.com/2014/Feb/08/serving-static-files-with-go.html

