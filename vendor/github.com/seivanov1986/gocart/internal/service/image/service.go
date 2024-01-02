package image

import (
	"github.com/seivanov1986/sql_client"

	"github.com/seivanov1986/gocart/internal/repository"
)

const (
	categoryType = 2
)

type Size struct {
	Height int
	Width  int
}

type service struct {
	hub       repository.Hub
	TrManager sql_client.TransactionManager
}

func New(hub repository.Hub, TrManager sql_client.TransactionManager) *service {
	return &service{hub: hub, TrManager: TrManager}
}
