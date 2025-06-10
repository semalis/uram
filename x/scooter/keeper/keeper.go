package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/semalis/uram/x/scooter/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
}

func NewKeeper(storeKey sdk.StoreKey) Keeper {
	return Keeper{storeKey: storeKey}
}

func (k Keeper) ScooterInfo(ctx context.Context, req *types.QueryScooterInfoRequest) (*types.QueryScooterInfoResponse, error) {
	return &types.QueryScooterInfoResponse{Info: "placeholder"}, nil
}
