package main

import g "github.com/jmdalida11/exploding-kittens/internal/game"

type Player struct {
	Id   string
	Name string
}

type Room struct {
	Name    string
	Game    *g.Game
	Players []Player
}

func CreateRoom(name string) *Room {
	return &Room{
		Name:    name,
		Game:    g.NewGame(),
		Players: []Player{},
	}
}
