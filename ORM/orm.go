package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
}

func initialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	checkError(err)
	defer db.Close()

	//User table created automatically and saves you creating the script in SQL to create the table.
	db.AutoMigrate(&User{})
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "allUsers Endpoint hit")

	db, err := gorm.Open("sqlite3", "test.db")
	checkError(err)
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User Endpoint hit")

	db, err := gorm.Open("sqlite3", "test.db")
	checkError(err)

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Succesfully created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Endpoint hit")

	db, err := gorm.Open("sqlite3", "test.db")
	checkError(err)
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User

	//Searches the db for the name.
	db.Where("name= ?", name).Find(&user)
	//Deletes user.
	db.Delete(&user)

	fmt.Fprintf(w, "User Deleted Successfully")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint hit")

	db, err := gorm.Open("sqlite3", "test.db")
	checkError(err)
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	//Searches the User from the name.
	db.Where("name = ?", name).Find(&user)

	//Change user.email as normal.
	user.Email = email

	//Save user in db.
	db.Save(&user)

	fmt.Fprintf(w, "Successfully Updated User")

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/users/{name}/{email}", newUser).Methods("POST")
	myRouter.HandleFunc("/users/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/users/{name}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {

	initialMigration()
	handleRequests()
}
