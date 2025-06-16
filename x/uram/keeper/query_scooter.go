package keeper

import (
	"context"
	"errors"

	"uram/x/uram/types"

	"cosmossdk.io/collections"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListScooter(ctx context.Context, req *types.QueryAllScooterRequest) (*types.QueryAllScooterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	scooters, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Scooter,
		req.Pagination,
		func(_ uint64, value types.Scooter) (types.Scooter, error) {
			return value, nil
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllScooterResponse{Scooter: scooters, Pagination: pageRes}, nil
}

func (q queryServer) GetScooter(ctx context.Context, req *types.QueryGetScooterRequest) (*types.QueryGetScooterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	scooter, err := q.k.Scooter.Get(ctx, req.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, sdkerrors.ErrKeyNotFound
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetScooterResponse{Scooter: scooter}, nil
}
