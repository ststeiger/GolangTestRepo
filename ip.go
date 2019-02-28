
package main

import (  
    // Standard library packages
    "fmt"
	"strconv"
	"log"
    "net"
	"net/http"
    "github.com/skratchdot/open-golang/open"
    // Third party packages
    "github.com/julienschmidt/httprouter"
)


/*
// FromRequest extracts the user IP address from req, if present.
func FromRequest(req *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}
	return userIP, nil
}
*/

// https://blog.golang.org/context/userip/userip.go
func getIP(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
    fmt.Fprintf(w, "<h1>static file server</h1><p><a href='./static'>folder</p></a>")
	
	ip, port, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		//return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
		
		fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		//return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
		fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
		return
	}
	
	forward := req.Header.Get("X-Forwarded-For")
	
	fmt.Fprintf(w, "<p>IP: %s</p>", ip)
	fmt.Fprintf(w, "<p>Port: %s</p>", port)
	fmt.Fprintf(w, "<p>Forwarded for: %s</p>", forward)
}



// go get github.com/toqueteos/webbrowser

func main() {  
    // Instantiate a new router
    r := httprouter.New()

    r.GET("/ip", getIP)

    // Add a handler on /test
    r.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        // Simply write some test data for now
        fmt.Fprint(w, "Welcome!\n")
    })

    //open.Run("https://google.com/")

    // open.Start("https://google.com")

/*
    // http://127.0.0.1:3000/test
    // Fire up the server
    http.ListenAndServe("localhost:3000", r)
    fmt.Println("ListenAndServe is blocking")  
    open.RunWith("http://localhost:3000/test", "firefox")  
    fmt.Println("Done")
*/
	
	myport := strconv.Itoa(10002);
	
	l, err := net.Listen("tcp", "localhost:" + myport)
	if err != nil {
	    log.Fatal(err)
	}
	
	// The browser can connect now because the listening socket is open.
	
	//err = open.Start("http://localhost:"+ myport + "/test")
	err = open.Start("http://localhost:"+ myport + "/ip")
	if err != nil {
	     log.Println(err)
	}
	
	// Start the blocking server loop.
	
	log.Fatal(http.Serve(l, r)) 
}