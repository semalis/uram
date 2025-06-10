package keeper

import (
	"context"
	"github.com/semalis/uram/x/scooter/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{Keeper: k}
}

func (s msgServer) RegisterScooter(ctx context.Context, msg *types.MsgRegisterScooter) (*types.MsgRegisterScooterResponse, error) {
	return &types.MsgRegisterScooterResponse{}, nil
}
