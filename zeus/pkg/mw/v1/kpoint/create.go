package kpoint

import (
	"context"
	"fmt"

	"github.com/Geapefurit/kline-back/zeus/pkg/db"
	"github.com/Geapefurit/kline-back/zeus/pkg/db/ent"
)

func (h *Handler) CreateKPointWithTx(ctx context.Context, tx *ent.Tx) error {
	sqlH := h.newSQLHandler()
	sql, err := sqlH.genCreateSQL()
	if err != nil {
		return err
	}
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return fmt.Errorf("fail create kpoint: %v", err)
	}
	return nil
}

func (h *Handler) CreateKPoint(ctx context.Context) error {
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		return h.CreateKPointWithTx(ctx, tx)
	})
}
