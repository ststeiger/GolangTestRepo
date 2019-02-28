// TestFile
package main

import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))
	// And here are some for slices and maps, which encode to JSON arrays and objects as you’d expect.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// You can use tags on struct field declarations to customize the encoded JSON key names. Check the definition of Response2 above to see an example of such tags.
	res2D := &Response2{
		Page:   1,
		Fruits: []string{"sandwich", "ice-cream", "sugar"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	fmt.Println("Hello World!")
}
