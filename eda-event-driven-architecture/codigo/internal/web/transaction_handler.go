package web

import (
	"encoding/json"
	"net/http"

	createtransaction "github.com/deirofelippe/curso-fullcycle/internal/usecase/create_transaction"
)

type WebTransactionHandler struct {
	CreateTransactionUsecase createtransaction.CreateTransactionUsecase
}

func NewWebTransactionHandler(createTransactionUsecase createtransaction.CreateTransactionUsecase) *WebTransactionHandler {
	return &WebTransactionHandler{
		CreateTransactionUsecase: createTransactionUsecase,
	}
}

func (h *WebTransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var dto createtransaction.CreateTransactionInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.CreateTransactionUsecase.Execute(dto)

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
