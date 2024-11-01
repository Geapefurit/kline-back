package tokenpair

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	npool "github.com/Geapefurit/kline-back/proto/kline/zeus/v1/tokenpair"

	"github.com/Geapefurit/kline-back/zeus/pkg/db"
	"github.com/Geapefurit/kline-back/zeus/pkg/db/ent"
	"github.com/Geapefurit/kline-back/zeus/pkg/db/ent/token"
	tokenpairent "github.com/Geapefurit/kline-back/zeus/pkg/db/ent/tokenpair"

	tokenpaircrud "github.com/Geapefurit/kline-back/zeus/pkg/crud/v1/tokenpair"
)

type queryHandler struct {
	*Handler
	stm   *ent.TokenPairSelect
	infos []*npool.TokenPair
	total uint32
}

func (h *queryHandler) selectTokenPair(stm *ent.TokenPairQuery) {
	h.stm = stm.Select(
		tokenpairent.FieldID,
		tokenpairent.FieldCreatedAt,
		tokenpairent.FieldUpdatedAt,
		tokenpairent.FieldTokenOneID,
		tokenpairent.FieldTokenTwoID,
		tokenpairent.FieldRemark,
	)
}

func (h *queryHandler) queryTokenPair(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.TokenPair.Query().Where(tokenpairent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(tokenpairent.ID(*h.ID))
	}
	h.selectTokenPair(stm)
	return nil
}

func (h *queryHandler) queryTokenPairs(ctx context.Context, cli *ent.Client) error {
	stm, err := tokenpaircrud.SetQueryConds(cli.TokenPair.Query(), h.Conds)
	if err != nil {
		return err
	}

	stmCount, err := tokenpaircrud.SetQueryConds(cli.TokenPair.Query(), h.Conds)
	if err != nil {
		return err
	}
	// stmCount.Modify(h.queryJoinToken)
	total, err := stmCount.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectTokenPair(stm)
	return nil
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(h.queryJoinToken)
}

func (h *queryHandler) queryJoinToken(s *sql.Selector) {
	tokenOneT := sql.Table(token.Table)
	tokenTwoT := sql.Table(token.Table)
	s.Join(tokenOneT).On(
		s.C(tokenpairent.FieldTokenOneID),
		tokenOneT.C(token.FieldID),
	).OnP(
		sql.EQ(tokenOneT.C(token.FieldDeletedAt), 0),
	).Join(tokenTwoT).On(
		s.C(tokenpairent.FieldTokenTwoID),
		tokenTwoT.C(token.FieldID),
	).OnP(
		sql.EQ(tokenTwoT.C(token.FieldDeletedAt), 0),
	).AppendSelect(
		sql.As(tokenOneT.C(token.FieldAddress), "token_one_address"),
		sql.As(tokenTwoT.C(token.FieldAddress), "token_two_address"),
	).Distinct()
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetTokenPair(ctx context.Context) (*npool.TokenPair, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTokenPair(cli); err != nil {
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

	return handler.infos[0], nil
}

func (h *Handler) GetTokenPairs(ctx context.Context) ([]*npool.TokenPair, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTokenPairs(ctx, cli); err != nil {
			return err
		}
		// handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(tokenpairent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
