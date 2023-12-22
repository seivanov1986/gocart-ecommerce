package image_to_product

import (
	"github.com/seivanov1986/sql_client"
)

const (
	limit = 8
)

type repository struct {
	db sql_client.DataBase
}

func New(db sql_client.DataBase) *repository {
	return &repository{
		db: db,
	}
}
