package yandex_feed

import (
	"context"
)

type Service interface {
	Generate(ctx context.Context, basePath, name, url, company string) string
}
