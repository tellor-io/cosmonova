package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSubmitValue{}, "cosmonova/SubmitValue", nil)
	cdc.RegisterConcrete(&MsgMakeRequest{}, "cosmonova/MakeRequest", nil)
	cdc.RegisterConcrete(&MsgRequestData{}, "cosmonova/RequestData", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitValue{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMakeRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestData{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
