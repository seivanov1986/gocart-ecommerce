package category

import (
	"context"
)

func (u *repository) DeleteIn(ctx context.Context, idCategories []int64) error {
	return u.db.DeleteIn(
		ctx,
		`DELETE FROM category WHERE id=?`,
		idCategories,
	)
}

func (u *repository) Delete(ctx context.Context, idCategory int64) error {
	_, err := u.db.ExecContext(
		ctx,
		`DELETE FROM category WHERE id=?`,
		idCategory,
	)
	return err
}
