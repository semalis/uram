package keeper

import (
	"context"
	"errors"
	"fmt"

	"uram/x/uram/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateScooter(ctx context.Context, msg *types.MsgCreateScooter) (*types.MsgCreateScooterResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	nextId, err := k.ScooterSeq.Next(ctx)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to get next id")
	}

	var scooter = types.Scooter{
		Id:       nextId,
		Creator:  msg.Creator,
		Location: msg.Location,
		Active:   msg.Active,
		Owner:    msg.Owner,
		Price:    msg.Price,
	}

	if err = k.Scooter.Set(
		ctx,
		nextId,
		scooter,
	); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set scooter")
	}

	return &types.MsgCreateScooterResponse{
		Id: nextId,
	}, nil
}

func (k msgServer) UpdateScooter(ctx context.Context, msg *types.MsgUpdateScooter) (*types.MsgUpdateScooterResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	var scooter = types.Scooter{
		Creator:  msg.Creator,
		Id:       msg.Id,
		Location: msg.Location,
		Active:   msg.Active,
		Owner:    msg.Owner,
		Price:    msg.Price,
	}

	// Checks that the element exists
	val, err := k.Scooter.Get(ctx, msg.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get scooter")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Scooter.Set(ctx, msg.Id, scooter); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update scooter")
	}

	return &types.MsgUpdateScooterResponse{}, nil
}

func (k msgServer) DeleteScooter(ctx context.Context, msg *types.MsgDeleteScooter) (*types.MsgDeleteScooterResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Checks that the element exists
	val, err := k.Scooter.Get(ctx, msg.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get scooter")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Scooter.Remove(ctx, msg.Id); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to delete scooter")
	}

	return &types.MsgDeleteScooterResponse{}, nil
}
