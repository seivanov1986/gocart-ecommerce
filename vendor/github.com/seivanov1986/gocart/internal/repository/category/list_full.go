package category

import (
	"context"
)

type CategoryListFullInput struct {
	Page     int64
	ParentID int64
}

type CategoryListFullOut struct {
	List  []CategoryListFullRow
	Total int64
}

type CategoryListFullRow struct {
	ID       int64  `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	ParentID int64  `db:"id_parent" json:"id_parent"`
}

func (i *repository) ListFull(ctx context.Context, in CategoryListFullInput) (*CategoryListFullOut, error) {
	pageRows := []CategoryListFullRow{}
	err := i.db.SelectContext(
		ctx,
		&pageRows,
		`SELECT id, name, id_parent FROM category LIMIT ?, ?`,
		in.ParentID, in.Page*limit, limit)
	if err != nil {
		return nil, err
	}

	var total int64
	err = i.db.GetContext(
		ctx,
		&total,
		`SELECT COUNT(*) FROM category`,
		in.ParentID,
	)
	if err != nil {
		return nil, err
	}

	return &CategoryListFullOut{
		List:  pageRows,
		Total: total,
	}, nil
}
