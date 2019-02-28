
package main


import (  
    // Standard library packages
    "fmt"
	"strings"
	"strconv"
	"log"
    "net"
	"net/http"
	"time"
	// "crypto/sha1"
	"crypto/sha256"
	"io"
	"math/rand"
	"encoding/hex"
	// Third party packages
    "github.com/julienschmidt/httprouter"
	"github.com/skratchdot/open-golang/open"
)


var mySingleSessionId int



// http://austingwalters.com/building-a-web-server-in-go-web-cookies/
func LoginCookie(username string) http.Cookie {
	// h := sha1.New()
	h := sha256.New()  
	mySingleSessionId = rand.Intn(100000000)
	
	//io.WriteString(h, username + strconv.Itoa( rand.Intn(100000000))  )
	io.WriteString(h, username + strconv.Itoa(mySingleSessionId)  )
	//fmt.Printf("% x", h.Sum(nil))
    
	// https://github.com/robertseaton/neptune/blob/master/pkgs/codify/codify.go
	// cookieValue := username + "=" + codify.SHA(username+strconv.Itoa(rand.Intn(100000000)))
	
	cookieValue := "username=" + username + "/sessionid=" + hex.EncodeToString(h.Sum(nil)) + // codify.SHA(username+strconv.Itoa(rand.Intn(100000000)))
	"/foo=abc"	
	expire := time.Now().AddDate(0, 0, 1)
	
	// &Cookie{Name: "cookie-9", Value: "expiring", Expires: time.Unix(1257894000, 0)},
	// &Cookie{Name: "cookie-2", Value: "two", MaxAge: 3600},
	// Cookie{Name: "cookie-3", Value: "three", Domain: ".example.com"},
	// Cookie{Name: "cookie-4", Value: "four", Path: "/restricted/"},
	return http.Cookie{Name: "SessionID", Value: cookieValue, Expires: expire, HttpOnly: true}
}


func lookupSessionID(username string)(string, string){
	if 1==1{
		h := sha256.New()  
		io.WriteString(h, username + strconv.Itoa(mySingleSessionId)  )
		return hex.EncodeToString(h.Sum(nil)), ""
	}
	
	return "", "I am an error" 
}



func setCookie(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
	cookie := LoginCookie("noob@gmail.com")
	
	fmt.Println("setting cookie")
	// http://golang.org/pkg/net/http/#SetCookie
	http.SetCookie(w, &cookie)
	fmt.Println("cookie set")
}

func getCookie(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
	isL := IsLoggedIn(req)
	// https://golang.org/pkg/fmt/
	fmt.Fprintf(w, "<p>IsLoggedIn: %t</p>", isL)
}





// https://coderwall.com/p/kjuyqw/get-environment-variables-as-a-map-in-golang
// http://stackoverflow.com/questions/12756782/go-http-post-and-use-cookies
func GetCookieValues(cookie *http.Cookie) map[string]string{
	var cookieValue = cookie.Value
	items := make(map[string]string)
	keyValuePairs := strings.Split(cookieValue, "/")
	
	for i:=0; i< len(keyValuePairs);i++{		
		kvp := strings.Split(keyValuePairs[i], "=")		
		items[kvp[0]] = kvp[1]
	}
	
	fmt.Println(items)
	return items
}

func IsLoggedIn(r *http.Request) bool {
	// https://golang.org/src/net/http/cookie_test.go
	// if g := req.Header.Get("Cookie"); g != tt.Raw {
	
	
    // Obtains cookie from users http.Request
	cookie, err := r.Cookie("SessionID")
	if err != nil {
		fmt.Println(err)
		return false
	}
	
	
	abc := GetCookieValues(cookie)
	
    // Split the sessionID to Username and ID (username+random)        
	email := abc["username"]
	sessionID := abc["sessionid"]
	
    fmt.Println("email: " + email)
	fmt.Println("sessionID: " + email)
	
    // Returns the expectedSessionID from the database
	expectedSessionID, errz := lookupSessionID(email)
	
	if errz != "" {
		fmt.Println(errz)
		return false
	}
	
	
	fmt.Println("expected: " + expectedSessionID)
	fmt.Println("actual: " + sessionID)
	
    // If SessionID matches the expected SessionID, it is Good
	if sessionID == expectedSessionID {
            // If you want to be really secure check IP
	    return true
	}
	
	return false
}



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
	
	// This will only be defined when site is accessed via non-anonymous proxy
	// and takes precedence over RemoteAddr
	forward := req.Header.Get("X-Forwarded-For")
	
	fmt.Fprintf(w, "<p>IP: %s</p>", ip)
	fmt.Fprintf(w, "<p>Port: %s</p>", port)
	fmt.Fprintf(w, "<p>Forwarded for: %s</p>", forward)
}


func main() {  
	myport := strconv.Itoa(10002);
	
    // Instantiate a new router
    r := httprouter.New()
	
    r.GET("/ip", getIP)
	
	
	r.GET("/setCookie", setCookie)
	r.GET("/getCookie", getCookie)
	
	
    // Add a handler on /test
    r.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        // Simply write some test data for now
        fmt.Fprint(w, "Welcome!\n")
    })	
	
	
	l, err := net.Listen("tcp", "localhost:" + myport)
	if err != nil {
	    log.Fatal(err)
	}
	// The browser can connect now because the listening socket is open.
	
	
	//err = open.Start("http://localhost:"+ myport + "/test")
	// err = open.Start("http://localhost:"+ myport + "/ip")
	err = open.RunWith("http://localhost:" + myport + "/setCookie", "chrome") 
	
	
	if err != nil {
	     log.Println(err)
	}
	
	// Start the blocking server loop.
	log.Fatal(http.Serve(l, r)) 
}
