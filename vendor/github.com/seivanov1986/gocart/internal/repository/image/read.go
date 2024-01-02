package image

import (
	"context"
)

type ImageReadRow struct {
	ID    int64  `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Path  string `db:"path" json:"path"`
	FType int64  `db:"ftype" json:"ftype"`
}

func (i *repository) Read(ctx context.Context, id int64) (*ImageReadRow, error) {
	row := ImageReadRow{}
	err := i.db.GetContext(
		ctx,
		&row,
		`SELECT id, name, path, ftype FROM image WHERE id = ?`,
		id)
	if err != nil {
		return nil, err
	}

	return &row, nil
}
