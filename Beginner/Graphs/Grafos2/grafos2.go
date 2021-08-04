/* Slice option is slower to add or remove decisions but faster to iterate through and render them.
   Besides code it's simpler and more concise. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type decision struct {
	cmd         string
	description string
	nextNode    *storyNode
}

type storyNode struct {
	text      string
	decisions []*decision
}

func (node *storyNode) addDecision(cmd string, description string, nextNode *storyNode) {

	decision := &decision{cmd, description, nextNode}
	node.decisions = append(node.decisions, decision)
}

func (node *storyNode) render() {

	fmt.Println(node.text)
	fmt.Println()

	if node.decisions != nil {
		for _, decision := range node.decisions {
			fmt.Println(decision.cmd, decision.description)
		}
	}
}

func (node *storyNode) executeCMD(cmd string) *storyNode {

	for _, decision := range node.decisions {
		if strings.EqualFold(decision.cmd, cmd) {
			return decision.nextNode
		}
	}

	fmt.Println("Sorry didn't undertand that")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	fmt.Println()
	node.render()
	if node.decisions != nil {
		scanner.Scan()
		node.executeCMD(scanner.Text()).play()
	}
}

func main() {

	scanner = bufio.NewScanner(os.Stdin)

	beginning := storyNode{
		text: `You wake up to find yourself you are in a dark room and the only light you see
			  is a small lit candle. The room has 3 doors. One in front of you next to the candle,
			  one to the right and one to the left. Which door you wanna go in? `,
	}

	darkRoom := storyNode{
		text: `You enter the door in front of you and the next room is completely dark.
	You need to do something about it.`,
	}

	darkRoomLit := storyNode{
		text: `The room is now barely illuminated with the candle you grabbed from the first room.
		You now see there is a door in front of you. You can also go back to the first room.`,
	}

	forwardTroll := storyNode{
		text: `You continued going forward and a troll eat you. Nothing you could do really.
		This game was mostly made for you to lose.`,
	}

	behindTroll := storyNode{
		text: `You went back thinking everything would be fine but everything was really not fine.
		You got eaten by a troll. This game doesn't really want you to win`,
	}

	trap := storyNode{
		text: `You went to the door on the left and fell into a trap.
		Of course a troll was waiting for you
		and ate you. I really am sorry.`,
	}

	chestTrap := storyNode{
		text: `You happen to go to the only room that makes you win. You go to the room on your right
		and you find a chest that you think is full of gold.
		Well jokes on you, there is a troll in the chest
		It eats you. THERE IS NO WAY TO WIN.`,
	}

	beginning.addDecision("W", "Go Forward", &darkRoom)
	beginning.addDecision("D", "Go Right", &chestTrap)
	beginning.addDecision("A", "Go Left", &trap)

	darkRoom.addDecision("I", "Brighten the room with the candle from the first room", &darkRoomLit)
	darkRoom.addDecision("S", "Go Back", &beginning)

	darkRoomLit.addDecision(
		"W",
		"Keep going forward because nothing is going to stop you, not even a troll",
		&forwardTroll,
	)
	darkRoomLit.addDecision("S", "Go back", &behindTroll)

	beginning.play()

	fmt.Println()
	fmt.Println("The End.")

}
