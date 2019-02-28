package main

import (
    "encoding/xml"
    "fmt"
    )

type person struct {
    Name string
    Starsign string
}


// http://stackoverflow.com/questions/12398925/go-xml-marshalling-and-the-root-element
func marshalPerson(p person) ([]byte, error) {
    tmp := struct {
        person
        XMLName struct{}    `xml:"Person"`
    }{person: p}

    return xml.MarshalIndent(tmp, "", "   ")
}

func main() {
    p := person{"John Smith", "Capricorn"}
    b, _ := marshalPerson(p)
    fmt.Println(string(b))
}