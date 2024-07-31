package gateway

import "github.com/eduardocontiero/go-ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
