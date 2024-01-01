package image

import (
	"context"
)

type ImageCreateInput struct {
	Name      string `db:"name"`
	ParentID  int64  `db:"id_parent"`
	Path      string `db:"path"`
	FType     int64  `db:"ftype"`
	CreatedAT int64  `db:"created_at"`
}

func (u *repository) Create(ctx context.Context, in ImageCreateInput) error {
	_, err := u.db.NamedExecContext(
		ctx, `
		INSERT INTO image (name, id_parent, path, ftype, created_at)
		VALUES (:name, :id_parent, :path, :ftype, :created_at)
	`, in)

	return err
}
