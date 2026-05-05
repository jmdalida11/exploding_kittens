package game

import "slices"

type Action struct {
	Action       string
	Card         Card
	Cards        []Card
	Value        int
	TargetPlayer string
}

type Player struct {
	Id         string
	Name       string
	Hands      []Card
	IsExploded bool
}

func (player *Player) AddCard(card Card) {
	player.Hands = append(player.Hands, card)
}

func (player *Player) HasCardOf(card Card) bool {
	return slices.Contains(player.Hands, card)
}

func (player *Player) CardTypeCount(card Card) int {
	count := 0
	for i := range player.Hands {
		if player.Hands[i] == card {
			count++
		}
	}
	return count
}

func (player *Player) RemoveCard(card Card) bool {
	for i := range player.Hands {
		if player.Hands[i] == card {
			player.Hands = slices.Delete(player.Hands, i, i+1)
			return true
		}
	}
	return false
}
