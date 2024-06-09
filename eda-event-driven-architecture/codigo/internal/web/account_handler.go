package web

import (
	"encoding/json"
	"net/http"

	createaccount "github.com/deirofelippe/curso-fullcycle/internal/usecase/create_account"
)

type WebAccountHandler struct {
	CreateAccountUsecase createaccount.CreateAccountUsecase
}

func NewWebAccountHandler(createAccountUsecase createaccount.CreateAccountUsecase) *WebAccountHandler {
	return &WebAccountHandler{
		CreateAccountUsecase: createAccountUsecase,
	}
}

func (h *WebAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var dto createaccount.CreateAccountInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.CreateAccountUsecase.Execute(dto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
