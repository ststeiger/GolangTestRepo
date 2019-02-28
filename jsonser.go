
package main

import (
    "encoding/json"
    "fmt"
)


type SearchResult struct {
    Date        string      `json:"date"`
    IdCompany   int         `json:"idCompany"`
    Company     string      `json:"company"`
    IdIndustry  interface{} `json:"idIndustry"`
    Industry    string      `json:"industry"`
    IdContinent interface{} `json:"idContinent"`
    Continent   string      `json:"continent"`
    IdCountry   interface{} `json:"idCountry"`
    Country     string      `json:"country"`
    IdState     interface{} `json:"idState"`
    State       string      `json:"state"`
    IdCity      interface{} `json:"idCity"`
    City        string      `json:"city"`
} //SearchResult


type SearchResults struct {
    NumberResults int            `json:"numberResults"`
    Results       []SearchResult `json:"results"`
} //type SearchResults


// http://stackoverflow.com/questions/17306358/golang-removing-fields-from-struct-or-hiding-them-in-json-response

func main() {
	msg := SearchResult{"a", 5, "Acme AG", "ACME_MwStNr.", "ACME", "EUR", "Europe", "CH", "Switzerland", "SG", "Sankt Gallen", "9450", "Altstaetten"}
	// I then encode and output the response like so:
	// err := json.NewEncoder(c.ResponseWriter).Encode(&msg)

	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
