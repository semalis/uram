package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgRegisterScooter{}

type MsgRegisterScooter struct {
	Creator   string
	ScooterId string
	Model     string
	Location  string
}

func (msg MsgRegisterScooter) Route() string { return RouterKey }
func (msg MsgRegisterScooter) Type() string  { return "RegisterScooter" }

func (msg MsgRegisterScooter) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{addr}
}

func (msg MsgRegisterScooter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgRegisterScooter) ValidateBasic() error {
	if msg.Creator == "" {
		return ErrInvalidCreator
	}
	return nil
}
