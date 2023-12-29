package attribute

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/attribute"
)

type AttributeSelectListIn struct {
	Query string
	Page  int64
}

func (u *service) SelectList(ctx context.Context, in AttributeSelectListIn) ([]attribute.AttributeSelectListRow, error) {
	return u.hub.Attribute().SelectList(ctx, attribute.AttributeSelectListInput{
		Query: in.Query,
		Page:  in.Page,
	})
}
