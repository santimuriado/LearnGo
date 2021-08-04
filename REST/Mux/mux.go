//updateMethod could be implemented.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

//homePage initialize.
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
	fmt.Println("Endpoint Hit: homePage")
}

//Initialize page which returns every article.
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticulos")
	json.NewEncoder(w).Encode(Articles)
}

//Loops every article that has the same id.
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleArticle")

	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}

}

// curl -H "Content-Type:application/json" -d '{"Id":"3", "Title":"Nuevo Post","desc":"Nueva
// Descripcion","content":"contenido nuevo"}' url.
func createNewArticle(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

// curl -X DELETE url.
func deleteArticle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

//Creates corresponding pages and the port they are going to use.
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {

	fmt.Println("REST API - MUX")
	Articles = []Article{
		{
			Id:      "1",
			Title:   "Generic title 1",
			Desc:    "Generic description 1",
			Content: "Generic content 1",
		},
		{
			Id:      "2",
			Title:   "Generic title 2",
			Desc:    "Generic description 2",
			Content: "Generic content 2",
		},
	}
	handleRequests()
}
