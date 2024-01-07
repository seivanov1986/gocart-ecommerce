package product

import (
	"context"
)

func (u *repository) DeleteIn(ctx context.Context, IDs []int64) error {
	return u.db.DeleteIn(
		ctx,
		`DELETE FROM product WHERE id IN (?)`,
		IDs,
	)
}

func (u *repository) Delete(ctx context.Context, ID int64) error {
	_, err := u.db.ExecContext(
		ctx,
		`DELETE FROM product WHERE id = ?`,
		ID,
	)

	return err
}
