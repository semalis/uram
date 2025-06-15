package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"cosmossdk.io/store/prefix"
	//storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/semalis/uram/x/scooter/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the QueryServer interface.
func NewQueryServerImpl(k Keeper) types.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

func (q queryServer) Scooter(ctx context.Context, req *types.QueryGetScooterRequest) (*types.QueryGetScooterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	//store := prefix.NewStore(q.k.storeService.OpenKVStore(ctx), []byte(types.ScooterKeyPrefix))

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(q.k.storeKey), types.KeyPrefix(types.ScooterKeyPrefix))
	//store := sdkCtx.KVStore(q.k.storeService.(storetypes.KVStoreKey))

	bz := store.Get(types.ScooterKey(req.Id))
	if bz == nil {
		return nil, status.Error(codes.NotFound, "scooter not found")
	}

	var scooter types.Scooter
	if err := q.k.cdc.Unmarshal(bz, &scooter); err != nil {
		return nil, status.Error(codes.Internal, "failed to unmarshal scooter")
	}

	return &types.QueryGetScooterResponse{Scooter: &scooter}, nil
}

func (q queryServer) ScooterAll(ctx context.Context, req *types.QueryAllScooterRequest) (*types.QueryAllScooterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	//store := prefix.NewStore(q.k.storeService.OpenKVStore(ctx), []byte(types.ScooterKeyPrefix))

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(q.k.storeKey), types.KeyPrefix(types.ScooterKeyPrefix))
	//store := sdkCtx.KVStore(q.k.storeService.(storetypes.KVStoreKey))

	var scooters []*types.Scooter

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var scooter types.Scooter
		if err := q.k.cdc.Unmarshal(value, &scooter); err != nil {
			return err
		}
		scooters = append(scooters, &scooter)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllScooterResponse{
		Scooters:   scooters,
		Pagination: pageRes,
	}, nil
}

func (q queryServer) Params(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	params, err := q.k.Params.Get(ctx)
	if err != nil && !errors.Is(err, collections.ErrNotFound) {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryParamsResponse{Params: params}, nil
}
