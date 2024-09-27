package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
    Name string
    Age int
}


var people []Person

func main() {
    http.HandleFunc("/people", peopleHandler)
    http.HandleFunc("/health", peopleHealthHandler)
    err := http.ListenAndServe("localhost:8080", nil)
    if err != nil { 
        log.Fatal(err)
    }

}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method { 
        case http.MethodGet:
            getPeople(w, r) 
        case http.MethodPost:
            postPerson(w, r)
        default:
            http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
    }
}


func getPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

func postPerson(w http.ResponseWriter, r *http.Request) {
    var person Person 
    err := json.NewDecoder(r.Body).Decode(&person)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    people = append(people, person)
    fmt.Println(w, "post new person: `%v`", person)
}

func peopleHealthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(w, "http web-server health")
}