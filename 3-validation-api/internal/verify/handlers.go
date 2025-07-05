package verify

import (
	"3-validation-api/config"
	requset "3-validation-api/pkg/request"
	"3-validation-api/pkg/utils"
	"encoding/json"
	"log"
	"net/http"
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

		requestSend, err := requset.HandleBody[RequestSend](&w, r)

		if err != nil {
			return
		}

		hash := utils.GenerateHash()
		DB = append(DB, RequestVerify{
			Email: requestSend.Email,
			Hash:  hash,
		})
		log.Println("Хэш добавлен в БД")

		err = utils.SendMail(requestSend.Email, hash, e.deps.config.Email, e.deps.config.Password, e.deps.config.Address)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
			return
		}

		log.Println("Email с верификацией отправлен")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Письмо с подтверждением отправлено"})
	}
}

func (e *EmailHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		hash := r.PathValue("hash")

		for idx, mark := range DB {
			if mark.Hash == hash {
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("true"))
				DB = utils.Remove(DB, idx)
				return
			}
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("false"))

	}
}
