package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Route struct {
	rooms []*Room
	mutex sync.Mutex
}

func CreateRoute() *Route {
	return &Route{
		rooms: []*Room{},
		mutex: sync.Mutex{},
	}
}

func (route *Route) InitRoutes() {
	http.HandleFunc("GET /get-rooms", route.getRooms)
	http.HandleFunc("POST /create-room", route.createRoom)
}

func (route *Route) parseBody(r *http.Request, out any) error {
	return json.NewDecoder(r.Body).Decode(out)
}

func (route *Route) createRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req CreateRoomRequest

	if err := route.parseBody(r, &req); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Invalid Room name", http.StatusBadRequest)
		return
	}

	route.mutex.Lock()
	defer route.mutex.Unlock()

	if route.isRoomNameExist(req.Name) {
		http.Error(w, "Room is already exist!", http.StatusBadRequest)
		return
	}
	route.rooms = append(route.rooms, CreateRoom(req.Name))

	json.NewEncoder(w).Encode(MessageResponse{
		Message: "Successfully created room " + req.Name,
	})
}

func (route *Route) isRoomNameExist(roomName string) bool {
	for i := range route.rooms {
		if route.rooms[i].Name == roomName {
			return true
		}
	}
	return false
}

func (route *Route) getRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	route.mutex.Lock()
	defer route.mutex.Unlock()

	var rooms []GetRoomsResponse = make([]GetRoomsResponse, 0)

	for _, room := range route.rooms {
		rooms = append(rooms, GetRoomsResponse{Name: room.Name})
	}

	json.NewEncoder(w).Encode(rooms)
}
