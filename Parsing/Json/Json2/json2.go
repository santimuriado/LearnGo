package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//This is the way to read the data without knowing the structure. Then it reads like a normal map.

func main() {

	jsonFile, err := os.Open("../users.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Opened correctly")

	defer jsonFile.Close()

	jsonBytes, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}

	json.Unmarshal([]byte(jsonBytes), &result)

	fmt.Println(result["users"])

}
