package gateway

import "github.com.br/deirofelippe/curso-fullcycle/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
