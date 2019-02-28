package main

import (
    "encoding/xml"
    "fmt"
    )

type person struct {
    Name string
    Starsign string
}

func main() {
    p := &person{"John Smith", "Capricorn"}
    b,_ := xml.MarshalIndent(p,"","   ")
    fmt.Println(string(b))
}