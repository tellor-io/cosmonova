package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestData = "request_data"

var _ sdk.Msg = &MsgRequestData{}

func NewMsgRequestData(creator string, qdata string, tip sdk.Coin) *MsgRequestData {
	return &MsgRequestData{
		Creator: creator,
		Qdata:   qdata,
		Tip:     tip,
	}
}

func (msg *MsgRequestData) Route() string {
	return RouterKey
}

func (msg *MsgRequestData) Type() string {
	return TypeMsgRequestData
}

func (msg *MsgRequestData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
