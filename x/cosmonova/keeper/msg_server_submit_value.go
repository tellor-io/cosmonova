package keeper

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	"cosmonova/x/cosmonova/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/crypto/sha3"
)

func (k msgServer) SubmitValue(goCtx context.Context, msg *types.MsgSubmitValue) (*types.MsgSubmitValueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		ctx.Logger().Error("invalid account address", "address", accAddr)
		return nil, err
	}
	val := sdk.ValAddress(accAddr)
	stakingKeeper := k.GetValidator(ctx, val)
	if !stakingKeeper {
		return nil, fmt.Errorf("validator is not bonded")
	}
	// reporter has to be validator
	var report = types.Report{
		QueryId:   msg.QueryId,
		QueryData: msg.QueryData,
		Nonce:     msg.Nonce,
		Value:     msg.Value,
		Timestamp: uint64(ctx.BlockTime().Unix()),
		Reporter:  msg.Creator,
	}
	if !compareQuery(msg.QueryId, msg.QueryData) {
		return nil, fmt.Errorf("invalid query data")
	}
	k.AddReport(ctx, report)
	event := sdk.NewEvent(
		"report submitted",
		sdk.NewAttribute("query_id", report.QueryId),
		sdk.NewAttribute("timestamp", string(report.Timestamp)),
	)

	// Emit the event
	ctx.EventManager().EmitEvent(event)
	return &types.MsgSubmitValueResponse{}, nil
}

func compareQuery(queryId, queryData string) bool {
	dataBytes, err := hex.DecodeString(queryData)
	if err != nil {
		return false
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write(dataBytes)
	resultHash := hash.Sum(nil)

	return strings.EqualFold(hex.EncodeToString(resultHash), queryId)
}
