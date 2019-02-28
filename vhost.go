
package main


import (  
    // Standard library packages
    "fmt"
    "io/ioutil"
    "net/http"
    //"time"
    "log"
    "github.com/julienschmidt/httprouter" 
    //"github.com/skratchdot/open-golang/open"
)


// Implement the ServerHTTP method on our new type
func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("srv for " + r.Host)
    // Check if a http.Handler is registered for the given host.
    // If yes, use it to handle the request.
    if handler := hs[r.Host]; handler != nil {
        handler.ServeHTTP(w, r)
    } else {
        // Handle host names for wich no handler is registered
        http.Error(w, "Forbidden", 403) // Or Redirect?
    }
}


func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Index!\n")
    var contbytes []byte
    contbytes, err  := ioutil.ReadFile("/root/Desktop/sorry_state.txt") 
	fmt.Println(err)
	fmt.Fprint(w, string(contbytes))
}

func Hello(w http.ResponseWriter, r *http.Request, a httprouter.Params) {
    fmt.Fprint(w, "Hello world!\n")
    fmt.Println(a)
}


type HostSwitch map[string]http.Handler

var hs HostSwitch;


func main() {
	// var content []byte, err error := ioutil.ReadFile("filename") 

    // Initialize a router as usual
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)
    
    // Make a new HostSwitch and insert the router (our http handler)
    // for example.com and port 12345
    hs := make(HostSwitch)
    //hs["example.com:12345"] = router
    hs["127.0.0.1:12345"] = router
    hs["localhost:12345"] = router
    
    // Use the HostSwitch to listen and serve on port 12345
    log.Fatal(http.ListenAndServe(":12345", hs))
}
