package main

type MessageResponse struct {
	Message string `json:"message"`
}

type CreateRoomRequest struct {
	Name string `json:"name" validate:"required"`
}

type GetRoomsResponse struct {
	Name string `json:"name"`
}
