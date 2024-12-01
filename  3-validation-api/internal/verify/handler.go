package verify

import (
	"github.com/SemenVologdin/go-advanced/pkg/req"
	"github.com/SemenVologdin/go-advanced/pkg/res"
	"log"
	"net/http"
)

type Handler struct{}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{}

	router.HandleFunc("POST /send", handler.Send(deps.SenderService, deps.HashService, deps.StorageService))
	router.HandleFunc("GET /verify/{hash}", handler.Verify(deps.StorageService))
}

func (handler *Handler) Send(sender Sender, hashService HashService, storage StorageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := req.HandleBody[SendRequest](r.Body)
		if err != nil {
			log.Printf("req.HandleBody: %v", err)
			errorResponse(w, "Ошибка при попытке получить тело запроса: "+err.Error(), http.StatusBadRequest)
			return
		}

		hash := hashService.HashString(request.Email)

		if err := storage.SaveEmailHash(r.Context(), request.Email, hash); err != nil {
			log.Printf("storage.SaveEmailHash: %v", err)
			errorResponse(w, "Ошибка при попытке сохранить hash емейла", http.StatusInternalServerError)
			return
		}

		if err := sender.Send(request.Email, hash); err != nil {
			log.Printf("sender.Send: %v", err)
			errorResponse(w, "Ошибка при попытке отправить письмо", http.StatusInternalServerError)
			return
		}

		if err := res.Json[struct{}](w, struct{}{}, http.StatusNoContent); err != nil {
			log.Printf("res.Json: %v", err)
			errorResponse(w, "Ошибка при попытке отправить ответ", http.StatusInternalServerError)
			return
		}
	}
}

func (handler *Handler) Verify(service StorageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		if hash == "" {
			errorResponse(w, "Параметр 'hash' не может быть пустым", http.StatusBadRequest)
			return
		}

		if _, err := service.EmailByHash(r.Context(), hash); err != nil {
			log.Printf("service.EmailByHash: %v", err)
			errorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := service.DeleteHash(r.Context(), hash); err != nil {
			log.Printf("service.DeleteHash: %v", err)
			errorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := res.Json[struct{}](w, struct{}{}, http.StatusNoContent); err != nil {
			log.Printf("res.Json: %v", err)
			errorResponse(w, "Ошибка при попытке отправить ответ", http.StatusInternalServerError)
			return
		}
	}
}
