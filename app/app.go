package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customer", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getCustomer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(rw, vars["customer_id"])
}

func createCustomer(rw http.ResponseWriter, r *http.Request) {

	fmt.Fprint(rw, "Post request received!")
}
