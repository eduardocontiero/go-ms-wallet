package gateway

import "github.com/eduardocontiero/go-ms-wallet/internal/entity"

type ClientGateway interface {
	FindByID(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
