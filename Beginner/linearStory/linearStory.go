package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyPage struct {
	text     string
	nextPage *storyPage //Linked list.
}

//Recursion
func playStory(page *storyPage) {
	//Base case
	if page == nil {
		return
	}
	fmt.Println(page.text)
	playStory(page.nextPage)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var input string
	fmt.Println("Testing scanner")
	//Waits for the user to press enter.
	scanner.Scan()
	fmt.Println("Any word here")
	fmt.Scanln(&input)
	fmt.Printf("The word written was: %v", input)
	scanner.Scan()

	page1 := storyPage{"The story begins", nil}
	page2 := storyPage{"The story continues", nil}
	page3 := storyPage{"Someone is in front of you", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3

	playStory(&page1)
}
