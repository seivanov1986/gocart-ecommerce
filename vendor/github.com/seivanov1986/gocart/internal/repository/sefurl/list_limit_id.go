package sefurl

import (
	"context"
)

type SefUrlListLimitIdRow struct {
	ID       int64   `db:"id" json:"id"`
	Url      string  `db:"url" json:"url"`
	Path     string  `db:"path" json:"path"`
	Name     string  `db:"name" json:"name"`
	Type     int64   `db:"type" json:"type"`
	IdObject int64   `db:"id_object" json:"id_object"`
	Template *string `db:"template" json:"template"`
}

func (i *repository) ListLimitId(ctx context.Context, offsetID int64) ([]SefUrlListLimitIdRow, error) {
	pageRows := []SefUrlListLimitIdRow{}
	err := i.db.SelectContext(
		ctx,
		&pageRows,
		`SELECT id, url, path, name, type, id_object, template FROM sefurl WHERE id > ? LIMIT ?`,
		offsetID, limit)
	if err != nil {
		return nil, err
	}

	return pageRows, nil
}
