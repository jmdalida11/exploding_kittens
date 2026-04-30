package game

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	EndMove                       = "end"
	ActivateCardMove              = "activate"
	BackExplodingKittenToDeckMove = "back_ek_to_deck"
	FavorMove                     = "favor"
	GiveCardMove                  = "give"
)

func ParseMove(input string) (*Action, error) {
	cmd := strings.Fields(input)
	cmdSize := len(cmd)

	if cmdSize == 0 {
		return nil, errors.New("No command entered")
	}

	action := cmd[0]

	switch action {
	case EndMove:
		return &Action{Action: action}, nil
	case ActivateCardMove:
		if cmdSize != 2 {
			return nil, errors.New(ActivateCardMove + " accepts 1 parameter: " + strings.Join(cmd, " "))
		}
		return &Action{Action: action, Card: Card(cmd[1])}, nil
	case BackExplodingKittenToDeckMove:
		if cmdSize != 2 {
			return nil, errors.New(BackExplodingKittenToDeckMove + " accepts 1 parameter: " + strings.Join(cmd, " "))
		}
		pos, err := strconv.Atoi(cmd[1])
		if err != nil {
			return nil, errors.New("Invalid input.")
		}
		return &Action{Action: action, Position: pos}, nil
	case FavorMove:
		if cmdSize != 2 {
			return nil, errors.New(FavorMove + " accepts 1 parameter (player name): " + strings.Join(cmd, " "))
		}
		return &Action{Action: action, TargetPlayer: cmd[1]}, nil
	case GiveCardMove:
		if cmdSize != 2 {
			return nil, errors.New(GiveCardMove + " accepts 1 parameter (card name): " + strings.Join(cmd, " "))
		}
		return &Action{Action: action, Card: Card(cmd[1])}, nil
	}

	return nil, fmt.Errorf("%s is not a valid action", action)
}
