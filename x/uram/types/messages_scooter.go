package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewMsgCreateScooter(creator string, location string, active bool, owner string, price sdk.Coin) *MsgCreateScooter {
	return &MsgCreateScooter{
		Creator:  creator,
		Location: location,
		Active:   active,
		Owner:    owner,
		Price:    &price,
	}
}

func NewMsgUpdateScooter(creator string, id uint64, location string, active bool, owner string, price sdk.Coin) *MsgUpdateScooter {
	return &MsgUpdateScooter{
		Id:       id,
		Creator:  creator,
		Location: location,
		Active:   active,
		Owner:    owner,
		Price:    &price,
	}
}

func NewMsgDeleteScooter(creator string, id uint64) *MsgDeleteScooter {
	return &MsgDeleteScooter{
		Id:      id,
		Creator: creator,
	}
}
