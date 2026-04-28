package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jmdalida11/exploding-kittens/internal/game"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	g := game.NewGame()

	g.AddPlayer("a", "A")
	g.AddPlayer("b", "B")
	g.AddPlayer("c", "C")

	g.Start()

	for true {
		currentPlayer := g.GetActivePlayer()

		switch g.GetState() {
		case game.PlayerMoveState:
			fmt.Println("\n== Player Cards ==")
			showPlayerCards(currentPlayer)
			fmt.Print("\nEnter move for player " + currentPlayer.Name + ": ")
		case game.BackExplodingKittenToDeckState:
			fmt.Printf("You Drawn Exploding Kitten! Choose where to put the Exploding Kitten card. (top of the deck) 0 to %d: ", g.GetDeckCount())
		case game.PlayerExploded:
			fmt.Printf("Player %s has been exploded!\n", currentPlayer.Name)
			g.SetState(game.PlayerMoveState)
			g.MoveToNextPlayer()

			if hasWinner, winner := g.HasWinner(); hasWinner {
				fmt.Println("\nGame Over!")
				fmt.Println("The winner is player " + winner.Name)
				return
			}

			continue
		}

		if scanner.Scan() {
			input := scanner.Text()

			if g.GetState() == game.BackExplodingKittenToDeckState {
				input = game.BackExplodingKittenToDeckMove + " " + input
			}

			action, err := game.ParseMove(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
			makeMove(g, *action)
		}
	}
}

func makeMove(g *game.Game, action game.Action) {
	switch g.GetState() {
	case game.PlayerMoveState:
		g.ActivePlayerMove(action)
	case game.BackExplodingKittenToDeckState:
		g.ReturnExplodingCardToDeck(action)
	}
}

func showPlayerCards(player *game.Player) {
	for _, card := range player.Hands {
		fmt.Println(card)
	}
}
