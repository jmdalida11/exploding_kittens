package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/jmdalida11/exploding-kittens/internal/game"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	g := game.NewGame()

	players := make(map[string]string)

	players["a"] = "A"
	players["b"] = "B"
	players["c"] = "C"

	for key, player := range players {
		g.AddPlayer(key, player)
	}

	g.Start()

	for true {
		currentPlayer := g.GetActivePlayer()

		switch g.GetState() {
		case game.PlayerMoveState:
			fmt.Println("\n== Player Cards ==")
			showPlayerCards(currentPlayer)
			fmt.Print("\nEnter move for player " + currentPlayer.Name + ": ")
		case game.BackExplodingKittenToDeckState:
			fmt.Printf("You've Drawn Exploding Kitten! Choose where to put the Exploding Kitten card. (top of the deck) 0 to %d: ", g.GetDeckCount())
		case game.FavorState:
			fmt.Printf("Enter player id: ")
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
		case game.SeeTheFutureState:
			fmt.Println()
			for i, card := range g.SeeTop3CardInDeck() {
				fmt.Println(strconv.Itoa(i+1) + ". " + string(card))
			}
			g.SetState(game.PlayerMoveState)
			continue
		case game.GiveCardState:
			fmt.Printf("\n%s", "Player "+g.GetTargetedPlayer().Name+" will give Player "+g.GetActivePlayer().Name+" a card.\n")
			fmt.Println("\n== Player Cards ==")
			showPlayerCards(g.GetTargetedPlayer())
			fmt.Print("(Player " + g.GetTargetedPlayer().Name + ") Enter card name to give: ")
		}

		if scanner.Scan() {
			input := scanner.Text()

			switch g.GetState() {
			case game.BackExplodingKittenToDeckState:
				input = game.BackExplodingKittenToDeckMove + " " + input
			case game.FavorState:
				input = game.FavorMove + " " + input
			case game.GiveCardState:
				input = game.GiveCardMove + " " + input
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
	case game.FavorState:
		if action.TargetPlayer != g.GetActivePlayer().Id && g.SetTargetedPlayer(action.TargetPlayer) {
			if len(g.GetTargetedPlayer().Hands) == 0 {
				g.SetState(game.PlayerMoveState)
			} else {
				g.SetState(game.GiveCardState)
			}
		} else {
			fmt.Println("Invalid Player Id")
		}
	case game.GiveCardState:
		targetPlayer := g.GetTargetedPlayer()
		if targetPlayer.RemoveCard(action.Card) {
			g.GetActivePlayer().AddCard(action.Card)
			g.SetState(game.PlayerMoveState)
		} else {
			fmt.Println("Invalid card " + action.Card)
		}
	}
}

func showPlayerCards(player *game.Player) {
	for _, card := range player.Hands {
		fmt.Println(card)
	}
}
