package gateway

import "github.com/deirofelippe/curso-fullcycle/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
	UpdateBalance(account *entity.Account) error
}
