package product

import (
	"context"
)

func (u *service) Delete(ctx context.Context, IDs []int64) error {
	for _, id := range IDs {
		u.TrManager.MakeTransaction(ctx, func(ctx context.Context) error {
			err := u.hub.Product().Delete(ctx, id)
			if err != nil {
				return err
			}

			err = u.hub.SefUrl().DeleteByObjectType(ctx, id, 3)
			if err != nil {
				return err
			}

			return nil
		})
	}

	return nil
}
