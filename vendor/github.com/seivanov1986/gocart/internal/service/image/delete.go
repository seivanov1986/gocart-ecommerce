package image

import (
	"context"
)

func (u *service) Delete(ctx context.Context, ID int64) error {
	return u.hub.Image().Delete(ctx, ID)
}
