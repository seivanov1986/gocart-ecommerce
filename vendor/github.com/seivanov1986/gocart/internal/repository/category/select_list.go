package category

import (
	"context"
)

type CategorySelectListInput struct {
	Query    string
	ParentID int64
}

type CategorySelectListOut struct {
	SelectList []CategorySelectListRow
	Total      int64
}

type CategorySelectListRow struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (i *repository) SelectList(ctx context.Context, in CategorySelectListInput) ([]CategorySelectListRow, error) {
	pageRows := []CategorySelectListRow{}
	err := i.db.SelectContext(
		ctx,
		&pageRows,
		`SELECT id, name FROM category WHERE name like ? LIMIT ?`,
		"%"+in.Query+"%", limit)
	return pageRows, err
}
