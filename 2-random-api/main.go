package main

import (
	"fmt"
	"net/http"
)

const PORT = "8080"

func main() {
	handler := NewHandler()

	if err := http.ListenAndServe(":"+PORT, handler.Init()); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
