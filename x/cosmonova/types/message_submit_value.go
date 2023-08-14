package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitValue = "submit_value"

var _ sdk.Msg = &MsgSubmitValue{}

func NewMsgSubmitValue(creator string, queryId string, queryData string, nonce uint64, value string) *MsgSubmitValue {
	return &MsgSubmitValue{
		Creator:   creator,
		QueryId:   queryId,
		QueryData: queryData,
		Nonce:     nonce,
		Value:     value,
	}
}

func (msg *MsgSubmitValue) Route() string {
	return RouterKey
}

func (msg *MsgSubmitValue) Type() string {
	return TypeMsgSubmitValue
}

func (msg *MsgSubmitValue) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitValue) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitValue) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
