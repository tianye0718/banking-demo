package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func greet(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello")
}
func getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Ye", City: "Auckland", Zipcode: "0931"},
		{Name: "Chang", City: "PalmersNorth", Zipcode: "3214"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(rw).Encode(customers)
	} else {
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(customers)
	}
}
