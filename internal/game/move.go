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
	ComboMove                     = "combo"
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
		return &Action{Action: action, Value: pos}, nil
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
	case ComboMove:
		if cmdSize < 2 {
			return nil, errors.New(ComboMove + " Invalid input: " + strings.Join(cmd, " "))
		}
		comboCount, err := strconv.Atoi(cmd[1])
		if err != nil {
			return nil, errors.New("Invalid input.")
		}
		vAction := &Action{Action: action, Value: comboCount}
		switch comboCount {
		case 2:
			if cmdSize != 4 {
				return nil, errors.New("Invalid combo")
			}
			vAction.Card = Card(cmd[2])
			vAction.TargetPlayer = cmd[3]
		case 3:
			if cmdSize != 5 {
				return nil, errors.New("Invalid combo")
			}
			vAction.Card = Card(cmd[2])
			vAction.TargetPlayer = cmd[3]
			vAction.Cards = []Card{Card(cmd[4])}
		case 5:
			if cmdSize != 8 {
				return nil, errors.New("Invalid combo")
			}
			vAction.Cards = []Card{Card(cmd[2]), Card(cmd[3]), Card(cmd[4]), Card(cmd[5]), Card(cmd[6])}
			vAction.Card = Card(cmd[7])
		default:
			return nil, errors.New("Invalid combo")
		}
		return vAction, nil
	}

	return nil, fmt.Errorf("%s is not a valid action", action)
}
