package page

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, in PageCreateInput) (*int64, error)
	Read(ctx context.Context, in PageReadInput) (*PageReadRow, error)
	Update(ctx context.Context, in PageUpdateInput) error
	DeleteIn(ctx context.Context, IDs []int64) error
	Delete(ctx context.Context, ID int64) error
	List(ctx context.Context, in PageListInput) (*PageListOut, error)
}
