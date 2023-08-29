package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMakeRequest = "make_request"

var _ sdk.Msg = &MsgMakeRequest{}

func NewMsgMakeRequest(creator string, queryData string, tip sdk.Coin) *MsgMakeRequest {
	return &MsgMakeRequest{
		Creator:   creator,
		QueryData: queryData,
		Tip:       tip,
	}
}

func (msg *MsgMakeRequest) Route() string {
	return RouterKey
}

func (msg *MsgMakeRequest) Type() string {
	return TypeMsgMakeRequest
}

func (msg *MsgMakeRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMakeRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMakeRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if !msg.Tip.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount is not a valid Coins object")
	}

	return nil
}
