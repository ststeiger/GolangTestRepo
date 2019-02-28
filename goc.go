package main

import (
    "C"
    "fmt"
)

//export PrintInt
func PrintInt(x int) {
    fmt.Println(x)
}

// http://stackoverflow.com/questions/32215509/using-go-code-in-an-existing-c-project
// go build -buildmode=c-archive goc.go
// go build -buildmode=c-shared goc.go 

// https://groups.google.com/forum/#!topic/golang-nuts/1oELh6joLQg
// Trying it on windows/amd64, looks like it isn't supported yet.  Is this planned for the 1.5 release? 
// It will not be in the 1.5 release.
// It would be nice if somebody worked on it for 1.6.
// https://golang.org/s/execmodes

// http://stackoverflow.com/questions/19431296/building-and-linking-dynamically-from-a-go-binary
// go build -linkshared hello.g
// go install -buildmode=shared std



func main() {
	fmt.Println("Hello world")
}
