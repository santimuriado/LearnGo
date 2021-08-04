package main

import (
	"bufio"
	"fmt"
	"os"
)

//Binary Tree.
type storyNode struct {
	text    string
	yesRoad *storyNode
	noRoad  *storyNode
}

func (nodo *storyNode) printStory(depth int) {

	//Adding indentation to see better the depth of the node.
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
	fmt.Println(nodo.text)
	if nodo.yesRoad != nil {
		nodo.yesRoad.printStory(depth + 1)
	}
	if nodo.noRoad != nil {
		nodo.noRoad.printStory(depth + 1)
	}
}

func (nodo *storyNode) play() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(nodo.text)

	if nodo.yesRoad != nil && nodo.noRoad != nil {

		for {
			scanner.Scan()
			response := scanner.Text()

			if response == "yes" {
				nodo.yesRoad.play()
				break
			} else if response == "no" {
				nodo.noRoad.play()
				break
			} else {
				fmt.Println("Only 'yes' and 'no' are accepted as responses.")
			}
		}
	}
}

func main() {

	root := storyNode{"You see a cave in front of you, do you want to enter?", nil, nil}
	win := storyNode{"You won", nil, nil}
	lose := storyNode{"You died", nil, nil}
	root.yesRoad = &lose
	root.noRoad = &win

	root.play()

	//Prints the story and the depth of each node.
	root.printStory(0)
}
