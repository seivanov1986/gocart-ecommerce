package image

import (
	"context"
)

type ImageListInput struct {
	Page     int64
	ParentID int64
}

type ImageListOut struct {
	List  []ImageListRow
	Total int64
}

type ImageListRow struct {
	ID    int64  `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Path  string `db:"path" json:"path"`
	FType int64  `db:"ftype" json:"ftype"`
}

const (
	limit = 8
)

func (i *repository) List(ctx context.Context, in ImageListInput) (*ImageListOut, error) {
	pageRows := []ImageListRow{}
	err := i.db.SelectContext(
		ctx,
		&pageRows,
		`SELECT id, name, path, ftype FROM image WHERE id_parent = ? LIMIT ?, ?`,
		in.ParentID, in.Page*limit, limit)
	if err != nil {
		return nil, err
	}

	var total int64
	err = i.db.GetContext(
		ctx,
		&total,
		`SELECT COUNT(*) FROM image WHERE id_parent = ?`,
		in.ParentID,
	)
	if err != nil {
		return nil, err
	}

	return &ImageListOut{
		List:  pageRows,
		Total: total,
	}, nil
}
