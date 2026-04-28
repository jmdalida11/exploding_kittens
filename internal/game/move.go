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
		return &Action{action: action}, nil
	case ActivateCardMove:
		if cmdSize != 2 {
			return nil, errors.New(ActivateCardMove + " accepts 1 parameter: " + strings.Join(cmd, " "))
		}
		return &Action{action: action, card: Card(cmd[1])}, nil
	case BackExplodingKittenToDeckMove:
		if cmdSize != 2 {
			return nil, errors.New(BackExplodingKittenToDeckMove + " accepts 1 parameter: " + strings.Join(cmd, " "))
		}
		pos, err := strconv.Atoi(cmd[1])
		if err != nil {
			return nil, errors.New("Invalid input.")
		}
		return &Action{action: action, position: pos}, nil
	}

	return nil, fmt.Errorf("%s is not a valid action", action)
}
