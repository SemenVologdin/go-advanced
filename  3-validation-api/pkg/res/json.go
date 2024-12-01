package res

import (
	"encoding/json"
	"net/http"
)

func Json[T any](writer http.ResponseWriter, data T, status int) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	return json.NewEncoder(writer).Encode(data)
}
