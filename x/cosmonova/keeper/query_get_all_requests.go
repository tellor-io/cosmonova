package keeper

import (
	"context"

	"cosmonova/x/cosmonova/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetAllRequests(goCtx context.Context, req *types.QueryGetAllRequestsRequest) (*types.QueryGetAllRequestsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var requests []types.DataRequest
	store := ctx.KVStore(k.storeKey)
	requestStore := prefix.NewStore(store, types.KeyPrefix(types.RequestKey))
	pageRes, err := query.Paginate(requestStore, req.Pagination, func(key []byte, value []byte) error {
		var request types.DataRequest
		if err := k.cdc.Unmarshal(value, &request); err != nil {
			return err
		}
		requests = append(requests, request)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	// TODO: Process the query
	_ = ctx

	return &types.QueryGetAllRequestsResponse{AllRequests: requests, Pagination: pageRes}, nil
}
