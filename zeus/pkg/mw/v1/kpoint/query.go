package kpoint

import (
	"context"
	"fmt"

	basetype "github.com/Geapefurit/kline-back/proto/kline/basetype/v1"
	kpointproto "github.com/Geapefurit/kline-back/proto/kline/zeus/v1/kpoint"

	"github.com/Geapefurit/kline-back/zeus/pkg/db"
	"github.com/Geapefurit/kline-back/zeus/pkg/db/ent"
	kpointent "github.com/Geapefurit/kline-back/zeus/pkg/db/ent/kpoint"

	kpointcrud "github.com/Geapefurit/kline-back/zeus/pkg/crud/v1/kpoint"
)

type queryHandler struct {
	*Handler
	stm   *ent.KPointSelect
	infos []*kpointproto.KPoint
	total uint32
}

func (h *queryHandler) selectKPoint(stm *ent.KPointQuery) {
	h.stm = stm.Select(
		kpointent.FieldID,
		kpointent.FieldCreatedAt,
		kpointent.FieldUpdatedAt,
		kpointent.FieldTokenPairID,
		kpointent.FieldKPointType,
		kpointent.FieldOpen,
		kpointent.FieldHigh,
		kpointent.FieldLow,
		kpointent.FieldClose,
		kpointent.FieldStartTime,
		kpointent.FieldEndTime,
	)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.KPointType = basetype.KPointType(basetype.KPointType_value[info.KPointTypeStr])
	}
}

func (h *queryHandler) queryKPoint(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.KPoint.Query().Where(kpointent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(kpointent.ID(*h.ID))
	}
	h.selectKPoint(stm)
	return nil
}

func (h *queryHandler) queryKPoints(ctx context.Context, cli *ent.Client) error {
	stm, err := kpointcrud.SetQueryConds(cli.KPoint.Query(), h.Conds)
	if err != nil {
		return err
	}

	stmCount, err := kpointcrud.SetQueryConds(cli.KPoint.Query(), h.Conds)
	if err != nil {
		return err
	}
	// stmCount.Modify(h.queryJoinToken)
	total, err := stmCount.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectKPoint(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetKPoint(ctx context.Context) (*kpointproto.KPoint, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryKPoint(cli); err != nil {
			return err
		}
		// handler.queryJoin()
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetKPoints(ctx context.Context) ([]*kpointproto.KPoint, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryKPoints(ctx, cli); err != nil {
			return err
		}
		// handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(kpointent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}

func (h *Handler) GetEarlistKPoints(ctx context.Context) ([]*kpointproto.KPoint, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryKPoints(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Asc(kpointent.FieldEndTime))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	return handler.infos, handler.total, nil
}

func (h *Handler) GetLatestKPoints(ctx context.Context) ([]*kpointproto.KPoint, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryKPoints(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(kpointent.FieldEndTime))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	return handler.infos, handler.total, nil
}
