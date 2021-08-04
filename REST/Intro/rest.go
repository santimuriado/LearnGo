package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNumber int            `json:"entry_number"`
	Species     PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {

	//Query
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	//Error checking.
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//Read from the byte stream.
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	//Region
	fmt.Println(responseObject.Name)
	//Amount of Pokemon.
	fmt.Println(len(responseObject.Pokemon))

	//Every pokemon.
	for i := 0; i < len(responseObject.Pokemon); i++ {
		fmt.Println(responseObject.Pokemon[i].EntryNumber)
		fmt.Println(responseObject.Pokemon[i].Species.Name)
	}

}
