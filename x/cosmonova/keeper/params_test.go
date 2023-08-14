package keeper_test

import (
	"testing"

	testkeeper "cosmonova/testutil/keeper"
	"cosmonova/x/cosmonova/types"

	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosanovaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
