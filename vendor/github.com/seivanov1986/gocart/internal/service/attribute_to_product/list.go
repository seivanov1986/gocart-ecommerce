package attribute_to_product

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/attribute_to_product"
)

type AttributeToProductListIn struct {
	ProductID int64
	Page      int64
}

func (c *service) List(ctx context.Context, in AttributeToProductListIn) (*attribute_to_product.AttributeToProductListOut, error) {
	return c.hub.AttributeToProduct().List(ctx, in.ProductID, in.Page)
}
