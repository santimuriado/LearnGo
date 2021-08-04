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
		text: `		Apareciste en una habitacion con la unica luz siendo una leve llama de una vela. 
		  Cada lado de la habitacion tiene una puerta negra que no tenes ni idea a donde va.
		  La puerta en frente tuyo esta totalmente nueva y sin usarse. Si la abris un poco te encontras con total oscuridad.
		  No te da buenas sensaciones.
		  La puerta a tu derecha esta un poco mas usada que la otra y ves que te lleva a una escalera que va para arriba.
		  Parece una buena opcion salvo que no sabes que puede haber arriba.
		  A la izquierda tenes la puerta en peor estado de todas. Si las otras parecian puertas pentagono esta es la tipica puerta de madera
		  que se cae a pedazos. A tener en cuenta que esto puede significar que es la puerta que mas se usa. O NO.`}

	darkRoom := storyNode{text: "Totalmente oscura, no podes ni verte la palma de la mano"}

	darkRoomLit := storyNode{
		text: `La habitacion ahora esta iluminada con la vela que te choreaste de la primera habitacion.
	  Podes seguir adelante o dar de baja todo, ser un cagon y volver pa atras.`,
	}

	forwardTroll := storyNode{
		text: `Y seguiste para adelante como un boludo y te comio un troll de la nada. 
		  Por que seguirias adelante y no volverias para atras no lo sabe ni Dios. Y si querias volver para atras tambien te comio un troll por tibio.`}

	behindTroll := storyNode{
		text: "Te comio el troll por ser un tibio que queres que te diga. Hubieras seguido para adelante."}

	trap := storyNode{
		text: `Fuiste por la puerta mas destruida del mundo porque te comiste el bait de que estaba usada por alguna buena razon. 
		  Bueno caiste en un trap pedazo de gil. Este juego no perdona.`}

	chest := storyNode{
		text: `Fuiste por la puerta de la derecha y te llevo por sorpresa. 
		  Llegaste a una habitacion con mas oro y joyerias que joyeria ricciardi y como son de tu confianza te agarraste todo.`}

	beginning.addDecision("W", "Ir para Adelante", &darkRoom)
	beginning.addDecision("D", "Ir a a la Derecha", &chest)
	beginning.addDecision("A", "Ir a la Izquierda", &trap)

	darkRoom.addDecision("I", "Iluminar la habitacion con la vela choreada", &darkRoomLit)
	darkRoom.addDecision("S", "Volver al beginning", &beginning)

	darkRoomLit.addDecision(
		"W",
		"Seguir para adelante porque tenes mas huevos que Marcos Rojo en el gol a Nigeria",
		&forwardTroll,
	)
	darkRoomLit.addDecision("S", "Volver para atras porque te comiste los mocos", &behindTroll)

	beginning.play()

	fmt.Println()
	fmt.Println("Termino bro")

}
