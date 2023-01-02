package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Pet struct {
	ID    string `json:"id"`
	Kind  string `json:"kind"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
	Owner *Owner `json:"owner"`
}

type Owner struct {
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}

var pets []Pet

func main() {
	fmt.Println("DOCKER-API-BASIC-GET")
	//seeding
	pets = append(pets,
		Pet{"1", "Cat Persia", "Momo", 2, &Owner{"Siska Stevani", "Malang"}},
		Pet{"2", "Owl", "Owly", 2, &Owner{"Anna", "Malang"}},
		Pet{"3", "Hamster", "Mimu", 2, &Owner{"Vanessa", "Surabaya"}},
		Pet{"4", "Marmut", "mickey", 1, &Owner{"Mickey", "Jakarta"}},
		Pet{"5", "Dog Labrador", "Haki", 2, &Owner{"Aris Antonius", "Sampit"}})

	//routing
	r := mux.NewRouter()
	r.HandleFunc("/pets", getAllPets).Methods("GET")
	r.HandleFunc("/pet/{id}", getPet).Methods("GET")

	//listen to a port
	log.Fatal(http.ListenAndServe(":8181", r))
}

func getAllPets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)

}

func getPet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params from request
	params := mux.Vars(r)

	//find id then return the response
	for _, pet := range pets {
		if pet.ID == params["id"] {
			json.NewEncoder(w).Encode(pet)
			return
		}
	}
	json.NewEncoder(w).Encode("Data not Found")
}
