package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

/* There's two ways to read the data. If the structure is known it's easier to make a predefined struct.
   The second way it's to unmarshal the JSON using a map[string] interface{} to parse the JSON into strings. */

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {

	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Opened correctly")

	defer jsonFile.Close()

	jsonBytes, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(jsonBytes, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User:Name " + users.Users[i].Name)
		fmt.Println("User:Age " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User:Type " + users.Users[i].Type)
		fmt.Println("Facebook Url " + users.Users[i].Social.Facebook)
		fmt.Println("Twitter Url " + users.Users[i].Social.Twitter)
	}

}
