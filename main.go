package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"prueba.jikkosoft/back/francisco/model"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome")
}

type Person struct {
	Name []int
}

func personCreate(w http.ResponseWriter, r *http.Request) {
	// Declare a new Person struct.
	var p []int
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Do something with the Person struct...
	model.Sort(p)
}

func sortArray(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	var sort model.ArrayToSort
	if err != nil {
		fmt.Fprintf(w, "not valid")
	}
	encjson, _ := json.Marshal(reqBody)
	json.Unmarshal(encjson, &sort)
	json.NewEncoder(w).Encode(sort)

}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/sortarray", sortArray).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
	fmt.Println("print consola")
}
