package handlers

import (
	"net/http"
	"purple/2-random-api/utils"
)

// Хендлер
type randomHandler struct{}

// Конструктор
func NewRandomHandler(mux *http.ServeMux) {
	handler := &randomHandler{}
	mux.HandleFunc("/number", handler.Number())
}

// Обработчик для поля
func (r *randomHandler) Number() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(utils.GetRandomNumber()))
	}

}
