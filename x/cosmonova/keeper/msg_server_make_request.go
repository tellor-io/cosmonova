package keeper

import (
	"context"
	"cosmonova/x/cosmonova/types"
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MakeRequest(goCtx context.Context, msg *types.MsgMakeRequest) (*types.MsgMakeRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	queryId, err := hex.DecodeString(msg.QueryData)
	if err != nil {
		return nil, err
	}
	tipper, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	err = k.bank.SendCoinsFromAccountToModule(ctx, tipper, types.ModuleName, sdk.NewCoins(msg.Tip))
	if err != nil {
		ctx.Logger().Error("error while sending coins to module", "error", err)
		return nil, err
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey))
	b := store.Get([]byte(queryId))
	var newRequest types.DataRequest
	if b == nil {
		newRequest = types.DataRequest{
			QueryData:    msg.QueryData,
			Value:        "",
			LastReported: 0,
			TipAmount:    msg.Tip,
			Status:       "pending",
		}
		store.Set([]byte(queryId), k.cdc.MustMarshal(&newRequest))
	} else {
		k.cdc.MustUnmarshal(b, &newRequest)
		newRequest.Status = "pending"
		newRequest.TipAmount = newRequest.TipAmount.Add(msg.Tip)
		store.Set([]byte(queryId), k.cdc.MustMarshal(&newRequest))
	}

	return &types.MsgMakeRequestResponse{}, nil
}
