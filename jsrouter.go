
package main


import (  
    // Standard library packages
    "fmt"
    "net/http"
    "time"
    "log"
    "github.com/julienschmidt/httprouter" 
    "github.com/skratchdot/open-golang/open"
)


// go get github.com/toqueteos/webbrowser
func main() {  
	// Instantiate a new router
	r := httprouter.New()

	// Add a handler on /test
	r.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Simply write some test data for now
		fmt.Fprint(w, "Welcome!\n")
	})
	
	
	
    go func() {
		for {
			time.Sleep(time.Second)
			
			// log.Println("Checking if started...")
			resp, err := http.Get("http://localhost:3000/test")
			if err != nil {
				log.Println("Failed:", err)
				continue
			}
			resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				log.Println("Not OK:", resp.StatusCode)
				continue
			}
			
			// Reached this point: server is up and running!
			break
        }
        
        log.Println("SERVER UP AND RUNNING!")
        open.RunWith("http://localhost:3000/test", "firefox")  
    }()
    
	// open.Run("https://google.com/")
	// open.Start("https://google.com")
	
	// http://127.0.0.1:3000/test
	// Fire up the server
	http.ListenAndServe("localhost:3000", r) 
	fmt.Println("hello")
}
