package keeper

import (
	"context"
	"encoding/hex"

	"cosmonova/x/cosmonova/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRequestData(goCtx context.Context, req *types.QueryGetRequestDataRequest) (*types.QueryGetRequestDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	qIdBytes, err := hex.DecodeString(req.QId)
	k.Logger(ctx).Error("qIdBytes", "qIdBytes", qIdBytes)
	if err != nil {
		return nil, err
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestDataKey))
	var request types.Requests
	b := store.Get(qIdBytes)
	if b == nil {
		return nil, nil
	}
	k.cdc.MustUnmarshal(b, &request)

	return &types.QueryGetRequestDataResponse{RequestData: &request}, nil
}
