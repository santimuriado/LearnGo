package main

import (
	"fmt"
)

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (page *storyPage) playStory() {
	if page == nil {
		return
	}
	fmt.Println(page.text)
	page.nextPage.playStory()

	// Tail call Elimination. Compiler turns it into a loop so the stack doesn't grow and it's more
	// efficient.
	//Either way in Golang the stack grows dinamically.
}

//Non recursive solution of playStory.
/* func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}
}*/

func (page *storyPage) addToEnd(text string) {
	for page.nextPage != nil {
		page = page.nextPage
	}
	page.nextPage = &storyPage{text, nil}
}

//O(1)
func (page *storyPage) addAfter(text string) {
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

func main() {

	page1 := storyPage{"You are in a forest", nil}
	page1.addToEnd("Walk straight")

	page1.addToEnd("You see a troll")
	page1.addAfter("Testing addAfter")
	page1.playStory()
}
