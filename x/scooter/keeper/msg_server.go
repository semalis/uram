package keeper

import (
	"context"

	"github.com/semalis/uram/x/scooter/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the MsgServer interface.
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{Keeper: k}
}

// Заглушки
func (m msgServer) CreateScooter(ctx context.Context, msg *types.MsgCreateScooter) (*types.MsgCreateScooterResponse, error) {
	return &types.MsgCreateScooterResponse{}, nil
}

func (m msgServer) UpdateScooter(ctx context.Context, msg *types.MsgUpdateScooter) (*types.MsgUpdateScooterResponse, error) {
	return &types.MsgUpdateScooterResponse{}, nil
}

func (m msgServer) DeleteScooter(ctx context.Context, msg *types.MsgDeleteScooter) (*types.MsgDeleteScooterResponse, error) {
	return &types.MsgDeleteScooterResponse{}, nil
}

/*
func (m msgServer) CreateScooter(goCtx context.Context, msg *types.MsgCreateScooter) (*types.MsgCreateScooterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(m.storeKey), []byte("scooter-"))

	scooter := types.Scooter{
		Id:       msg.Id,
		Owner:    msg.Creator,
		Location: msg.Location,
		Active:   true,
		Price:    msg.Price,
	}

	bz, err := m.cdc.Marshal(&scooter)
	if err != nil {
		return nil, err
	}

	key := sdk.Uint64ToBigEndian(msg.Id)
	store.Set(key, bz)

	return &types.MsgCreateScooterResponse{}, nil
}
*/

//func (m msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
//	return &types.MsgUpdateParamsResponse{}, nil
//}

//func (m msgServer) CreateScooter(goCtx context.Context, msg *types.MsgCreateScooter) (*types.MsgCreateScooterResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//        store := prefix.NewStore(corecompat.KVStoreAdapter(m.Keeper.storeService.OpenKVStore(ctx)), []byte(types.ScooterKey))
//	_ = store
//	return &types.MsgCreateScooterResponse{}, nil
//}

//func (m msgServer) UpdateScooter(goCtx context.Context, msg *types.MsgUpdateScooter) (*types.MsgUpdateScooterResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//        store := prefix.NewStore(corecompat.KVStoreAdapter(m.Keeper.storeService.OpenKVStore(ctx)), []byte(types.ScooterKey))
//	_ = store
//	return &types.MsgUpdateScooterResponse{}, nil
//}

//func (m msgServer) DeleteScooter(goCtx context.Context, msg *types.MsgDeleteScooter) (*types.MsgDeleteScooterResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//        store := prefix.NewStore(corecompat.KVStoreAdapter(m.Keeper.storeService.OpenKVStore(ctx)), []byte(types.ScooterKey))
//	_ = store
//	return &types.MsgDeleteScooterResponse{}, nil
//}
