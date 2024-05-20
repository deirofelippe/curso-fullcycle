package gateway

import "github.com.br/deirofelippe/curso-fullcycle/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
}
