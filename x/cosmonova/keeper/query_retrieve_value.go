package keeper

import (
	"context"
	"fmt"

	"cosmonova/x/cosmonova/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RetrieveValue(goCtx context.Context, req *types.QueryRetrieveValueRequest) (*types.QueryRetrieveValueResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	report, found := k.GetReport(ctx, req.QueryId, req.Timestamp)

	if !found {
		fmt.Println("can't find report")
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryRetrieveValueResponse{Report: &report}, nil
}
