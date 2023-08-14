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

func (k Keeper) RetrieveValueAll(goCtx context.Context, req *types.QueryRetrieveValueAllRequest) (*types.QueryRetrieveValueAllResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	var reports []types.Report
	store := ctx.KVStore(k.storeKey)
	reportStore := prefix.NewStore(store, types.KeyPrefix(types.ReportKey))
	pageRes, err := query.Paginate(reportStore, req.Pagination, func(key []byte, value []byte) error {
		var report types.Report
		if err := k.cdc.Unmarshal(value, &report); err != nil {
			return err
		}
		reports = append(reports, report)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryRetrieveValueAllResponse{Report: reports, Pagination: pageRes}, nil
}
