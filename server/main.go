package main

import (
	"fmt"
	"net/http"
)

const PORT = 3000

func main() {
	fmt.Printf("Server running on :%d", PORT)

	CreateRoute().InitRoutes()

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
