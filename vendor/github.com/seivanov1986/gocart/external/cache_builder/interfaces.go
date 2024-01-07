package cache_builder

import (
	"context"
)

type Service interface {
	RenderObject(ctx context.Context, url string) ([]byte, error)
}
