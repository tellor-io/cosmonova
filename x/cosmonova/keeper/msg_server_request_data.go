package keeper

import (
	"context"
	"encoding/hex"

	"cosmonova/x/cosmonova/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/crypto/sha3"
)

func (k msgServer) RequestData(goCtx context.Context, msg *types.MsgRequestData) (*types.MsgRequestDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	dataBytes, err := hex.DecodeString(msg.Qdata)
	if err != nil {
		return nil, err
	}
	hash := sha3.NewLegacyKeccak256()
	hash.Write(dataBytes)
	qID := hash.Sum(nil)

	tipper, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestDataKey))

	b := store.Get(qID)
	var newRequest types.Requests
	if b == nil {
		newRequest = types.Requests{
			QueryData:     msg.Qdata,
			ValueList:     []*types.Values{},
			Tip:           msg.Tip,
			TimeofRequest: ctx.BlockTime().Unix(),
			Status:        "pending",
		}
	} else {
		k.cdc.MustUnmarshal(b, &newRequest)
		if newRequest.Status == "pending" {
			newRequest.Tip = newRequest.Tip.Add(msg.Tip)
		} else {
			newRequest.Tip = newRequest.Tip.Add(msg.Tip)
			newRequest.Status = "pending"
			newRequest.TimeofRequest = ctx.BlockTime().Unix()
		}
	}
	err = k.bank.SendCoinsFromAccountToModule(ctx, tipper, types.ModuleName, sdk.NewCoins(msg.Tip))
	if err != nil {
		ctx.Logger().Error("error while sending coins to module", "error", err)
		return nil, err
	}
	store.Set(qID, k.cdc.MustMarshal(&newRequest))
	return &types.MsgRequestDataResponse{}, nil
}
