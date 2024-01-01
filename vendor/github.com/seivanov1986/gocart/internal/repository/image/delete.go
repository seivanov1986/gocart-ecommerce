package image

import (
	"context"
)

func (u *repository) Delete(ctx context.Context, idImage int64) error {
	_, err := u.db.ExecContext(
		ctx,
		`DELETE FROM image WHERE id=?`,
		idImage,
	)
	return err
}
