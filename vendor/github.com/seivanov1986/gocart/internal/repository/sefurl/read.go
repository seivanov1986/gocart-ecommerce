package sefurl

import (
	"context"
)

type SefUrlReadInput struct {
	ID int64
}

type SefUrlReadRow struct {
	ID       int64   `db:"id" json:"id"`
	Url      string  `db:"url" json:"url"`
	Path     string  `db:"path" json:"path"`
	Name     string  `db:"name" json:"name"`
	Type     int64   `db:"type" json:"type"`
	ObjectID int64   `db:"id_object" json:"id_object"`
	Template *string `db:"template" json:"template"`
}

func (i *repository) Read(ctx context.Context, url string) (*SefUrlReadRow, error) {
	row := SefUrlReadRow{}
	err := i.db.GetContext(
		ctx,
		&row,
		`SELECT id, url, path, name, type, id_object, template FROM sefurl WHERE url = ?`,
		url)
	if err != nil {
		return nil, err
	}

	return &row, nil
}
