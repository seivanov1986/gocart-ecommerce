package category

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/category"
)

type CategorySelectListIn struct {
	Query    string
	ParentID int64
}

func (u *service) SelectList(ctx context.Context, in CategorySelectListIn) ([]category.CategorySelectListRow, error) {
	return u.hub.Category().SelectList(ctx, category.CategorySelectListInput{
		Query:    in.Query,
		ParentID: in.ParentID,
	})
}
