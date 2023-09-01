package keeper

import (
	"context"

	"cosmonova/x/cosmonova/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRequestCount(goCtx context.Context, req *types.QueryGetRequestCountRequest) (*types.QueryGetRequestCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryGetRequestCountResponse{}, nil
}
