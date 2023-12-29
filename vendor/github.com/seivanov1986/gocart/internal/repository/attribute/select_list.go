package attribute

import (
	"context"
)

type AttributeSelectListInput struct {
	Query string
	Page  int64
}

type AttributeSelectListRow struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (i *repository) SelectList(ctx context.Context, in AttributeSelectListInput) ([]AttributeSelectListRow, error) {
	pageRows := []AttributeSelectListRow{}
	err := i.db.SelectContext(
		ctx,
		&pageRows,
		`SELECT id, name FROM attribute WHERE name like ? LIMIT ?`,
		"%"+in.Query+"%", limit)
	return pageRows, err
}
