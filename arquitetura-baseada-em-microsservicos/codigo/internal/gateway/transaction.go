package gateway

import "github.com.br/deirofelippe/curso-fullcycle/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
