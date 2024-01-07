package category

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, in CategoryCreateInput) (*int64, error)
	Read(ctx context.Context, categoryID int64) (*CategoryReadRow, error)
	Update(ctx context.Context, in CategoryUpdateInput) error
	DeleteIn(ctx context.Context, idCategories []int64) error
	Delete(ctx context.Context, idCategory int64) error
	List(ctx context.Context, in CategoryListInput) (*CategoryListOut, error)

	SelectList(ctx context.Context, in CategorySelectListInput) ([]CategorySelectListRow, error)
	ListFull(ctx context.Context, in CategoryListFullInput) (*CategoryListFullOut, error)
}
