package cosmonova_test

import (
	"testing"

	keepertest "cosmonova/testutil/keeper"
	"cosmonova/testutil/nullify"
	"cosmonova/x/cosmonova"
	"cosmonova/x/cosmonova/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmonovaKeeper(t)
	cosmonova.InitGenesis(ctx, *k, genesisState)
	got := cosmonova.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
