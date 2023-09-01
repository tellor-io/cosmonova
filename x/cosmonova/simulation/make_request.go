package simulation

import (
	"math/rand"

	"cosmonova/x/cosmonova/keeper"
	"cosmonova/x/cosmonova/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgMakeRequest(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMakeRequest{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the MakeRequest simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "MakeRequest simulation not implemented"), nil, nil
	}
}
