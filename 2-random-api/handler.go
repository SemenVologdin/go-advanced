package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
)

const (
	MIN = 1
	MAX = 6
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/random", random(MIN, MAX))

	return router
}

func random(min, max int) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		number := rand.IntN(max+1-min) + min
		if _, err := writer.Write([]byte(strconv.Itoa(number))); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}
	}
}
