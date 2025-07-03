package email

import (
	"encoding/json"
	"log"
	"net/http"
	"purple/validation/config"
	"purple/validation/pkg/utils"
	"slices"
	"strings"
)

type EmailHandler struct {
	deps EmailDeps
}

type EmailDeps struct {
	config *config.EmailConfig
}

func NewEmailHandler(mux *http.ServeMux, config *config.Config) {
	handler := &EmailHandler{
		deps: EmailDeps{
			config: config.Email,
		},
	}
	mux.HandleFunc("POST /send", handler.Send())
	mux.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (e *EmailHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req struct {
			Email string `json:"email"`
		}

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(w, "Ошибка при чтении тела запроса", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		hash := utils.GenerateHash()
		DB = append(DB, hash)
		log.Println("Хэш добавлен в БД")

		err = utils.SendMail(req.Email, hash, e.deps.config.Email, e.deps.config.Password, e.deps.config.Address)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
		}

		log.Println("Email с верификацией отправлен")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Письмо с подтверждением отправлено"})
	}
}

func (e *EmailHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		parts := strings.Split(path, "/")

		if len(parts) < 2 {
			http.NotFound(w, r)
			return
		}

		hash := parts[len(parts)-1]

		if slices.Contains(DB, hash) {
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(map[string]string{"message": "Welcome"})
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "error"})
		}
	}
}
