//nolint:nolintlint,dupl
package kpoint

import (
	"context"

	kpointproto "github.com/Geapefurit/kline-back/proto/kline/zeus/v1/kpoint"
	kpoint "github.com/Geapefurit/kline-back/zeus/pkg/mw/v1/kpoint"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetKPoint(ctx context.Context, in *kpointproto.GetKPointRequest) (*kpointproto.GetKPointResponse, error) {
	handler, err := kpoint.NewHandler(
		ctx,
		kpoint.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetKPoint",
			"In", in,
			"Error", err,
		)
		return &kpointproto.GetKPointResponse{}, status.Error(codes.Internal, "internal server err")
	}

	info, err := handler.GetKPoint(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetKPoint",
			"In", in,
			"Error", err,
		)
		return &kpointproto.GetKPointResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &kpointproto.GetKPointResponse{
		Info: info,
	}, nil
}

func (s *Server) GetKPoints(ctx context.Context, in *kpointproto.GetKPointsRequest) (*kpointproto.GetKPointsResponse, error) {
	handler, err := kpoint.NewHandler(
		ctx,
		kpoint.WithConds(in.Conds),
		kpoint.WithOffset(in.GetOffset()),
		kpoint.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetKPoints",
			"In", in,
			"Error", err,
		)
		return &kpointproto.GetKPointsResponse{}, status.Error(codes.Internal, "internal server err")
	}

	infos, total, err := handler.GetKPoints(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetKPoints",
			"In", in,
			"Error", err,
		)
		return &kpointproto.GetKPointsResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &kpointproto.GetKPointsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
