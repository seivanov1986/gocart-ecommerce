package sefurl

import (
	"context"

	"github.com/seivanov1986/gocart/internal/repository/sefurl"
)

type Service interface {
	List(ctx context.Context, in SefUrlListIn) (*sefurl.SefUrlListOut, error)
}
