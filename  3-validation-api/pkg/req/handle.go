package req

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"io"
	"log"
)

func HandleBody[T any](body io.ReadCloser) (*T, error) {
	var globalError error

	var request T
	if err := json.NewDecoder(body).Decode(&request); err != nil {
		log.Printf("json.NewDecoder.Decode %+v\n", err)
		return nil, err
	}
	defer func() {
		if err := body.Close(); err != nil {
			globalError = err
		}
	}()

	if err := validator.New().Struct(request); err != nil {
		return nil, err
	}

	return &request, globalError
}
