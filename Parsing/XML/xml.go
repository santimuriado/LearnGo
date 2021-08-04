package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func main() {

	XMLFile, err := os.Open("users.xml")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Correctly opened users.xml")

	defer XMLFile.Close()

	xmlBytes, _ := ioutil.ReadAll(XMLFile)

	var users Users

	xml.Unmarshal(xmlBytes, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
		fmt.Println("Twitter Url: " + users.Users[i].Social.Twitter)
	}
}
