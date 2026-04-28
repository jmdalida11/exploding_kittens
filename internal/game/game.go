package game

import (
	"fmt"
	"slices"
)

type GameState int

const (
	PlayerMoveState = iota
	BackExplodingKittenToDeckState
	PlayerExploded
	CardEffectState
)

type Game struct {
	deck              []Card
	discardPile       []Card
	players           map[string]*Player
	playerPositionId  []string
	activePlayerIndex int
	state             GameState
}

func NewGame() *Game {
	return &Game{
		deck:              []Card{},
		discardPile:       []Card{},
		players:           make(map[string]*Player),
		activePlayerIndex: 0,
		state:             PlayerMoveState,
	}
}

func popLastCard(cards []Card) (Card, []Card) {
	last := cards[len(cards)-1]
	cards = cards[:len(cards)-1]
	return last, cards
}

func (game *Game) generateDeckCard(card Card, count int) {
	for range count {
		game.deck = append(game.deck, card)
	}
}

func (game *Game) generateDeckCards() {
	game.generateDeckCard(Defuse, 1)
	game.generateDeckCard(Nope, 5)
	game.generateDeckCard(SeeTheFuture, 5)
	game.generateDeckCard(Attack, 4)
	game.generateDeckCard(Skip, 4)
	game.generateDeckCard(Favor, 4)
	game.generateDeckCard(Shuffle, 4)
	game.generateDeckCard(Tacocat, 4)
	game.generateDeckCard(Cattermelon, 4)
	game.generateDeckCard(HairyPotatoCat, 4)
	game.generateDeckCard(BeardCat, 4)
	game.generateDeckCard(RainbowRalphingCat, 4)
	randomizer(game.deck)
}

func (game *Game) GetState() GameState {
	return game.state
}

func (game *Game) SetState(state GameState) {
	game.state = state
}

func (game *Game) GetDeckCount() int {
	return len(game.deck)
}

func (game *Game) AddPlayer(id string, name string) {
	if _, exist := game.players[id]; exist {
		fmt.Println("Id % is already exist", id)
		return
	}
	game.players[id] = &Player{Id: id, Name: name, Hands: []Card{}, IsExploded: false}
	game.playerPositionId = append(game.playerPositionId, id)
}

func (game *Game) RemovePlayer(id string) {
	delete(game.players, id)
}

func (game *Game) DrawCardFromDeck() Card {
	card, deck := popLastCard(game.deck)
	game.deck = deck
	return card
}

func (game *Game) GetActivePlayer() *Player {
	return game.players[game.playerPositionId[game.activePlayerIndex]]
}

func (game *Game) DrawActivePlayerCard() Card {
	activePlayer := game.GetActivePlayer()
	cardDrawn := game.DrawCardFromDeck()
	activePlayer.Hands = append(activePlayer.Hands, cardDrawn)
	return cardDrawn
}

func (game *Game) generatePlayerCards() {
	for _, player := range game.players {
		player.Hands = append(player.Hands, Defuse)
	}
	for range 4 {
		for _, player := range game.players {
			player.Hands = append(player.Hands, game.DrawCardFromDeck())
		}
	}
}

func (game *Game) MoveToNextPlayer() {
	for true {
		game.activePlayerIndex += 1
		if game.activePlayerIndex >= len(game.playerPositionId) {
			game.activePlayerIndex = 0
		}
		if !game.GetActivePlayer().IsExploded {
			break
		}
	}
}

func (game *Game) HasWinner() (bool, *Player) {
	var winner *Player = nil
	for _, player := range game.players {
		if !player.IsExploded {
			if winner != nil {
				return false, nil
			}
			winner = player
		}
	}
	return true, winner
}

func (game *Game) Start() {
	fmt.Print("\nWelcome to the Exploding Kittens Game!\n")

	if len(game.players) < 2 {
		fmt.Println("Should at least have 2 players to start the game.")
		return
	}

	randomizer(game.playerPositionId)
	game.generateDeckCards()
	game.generatePlayerCards()
	game.generateDeckCard(ExplodingKitten, len(game.players)-1)
	randomizer(game.deck)
}

func (game *Game) ReturnExplodingCardToDeck(action Action) {
	deckSize := len(game.deck)
	if action.position >= 0 && action.position <= deckSize {
		game.deck = slices.Insert(game.deck, deckSize-action.position, ExplodingKitten)
		game.MoveToNextPlayer()
		game.state = PlayerMoveState
	} else {
		fmt.Println("Invalid deck position")
	}
}

func (game *Game) ActivePlayerMove(action Action) {
	if game.state != PlayerMoveState {
		fmt.Println("Not a valid state to do the current action")
		return
	}

	activePlayer := game.GetActivePlayer()

	switch action.action {
	case EndMove:
		cardDrawn := game.DrawActivePlayerCard()
		if cardDrawn == ExplodingKitten {
			if activePlayer.RemoveCard(Defuse) {
				game.state = BackExplodingKittenToDeckState
				activePlayer.RemoveCard(ExplodingKitten)
			} else {
				activePlayer.IsExploded = true
				for _, card := range activePlayer.Hands {
					game.discardPile = append(game.discardPile, card)
				}
				activePlayer.Hands = []Card{}
				game.state = PlayerExploded
			}
		} else {
			game.MoveToNextPlayer()
		}
	case ActivateCardMove:
		card := action.card
		player := game.GetActivePlayer()
		if game.GetActivePlayer().RemoveCard(card) {
			CardEffect(card, game)
		} else {
			fmt.Printf("Player %s has no card %s", player.Name, card)
		}
	default:
		fmt.Println("Invalid action: " + action.action)
	}
}
