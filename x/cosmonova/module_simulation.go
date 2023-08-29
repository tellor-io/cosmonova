package cosmonova

import (
	"math/rand"

	"cosmonova/testutil/sample"
	cosmonovasimulation "cosmonova/x/cosmonova/simulation"
	"cosmonova/x/cosmonova/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = cosmonovasimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgSubmitValue = "op_weight_msg_submit_value"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitValue int = 100

	opWeightMsgMakeRequest = "op_weight_msg_make_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMakeRequest int = 100

	opWeightMsgRequestData = "op_weight_msg_request_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRequestData int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cosanovaGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cosanovaGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSubmitValue int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitValue, &weightMsgSubmitValue, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitValue = defaultWeightMsgSubmitValue
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitValue,
		cosmonovasimulation.SimulateMsgSubmitValue(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMakeRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMakeRequest, &weightMsgMakeRequest, nil,
		func(_ *rand.Rand) {
			weightMsgMakeRequest = defaultWeightMsgMakeRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMakeRequest,
		cosmonovasimulation.SimulateMsgMakeRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRequestData int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRequestData, &weightMsgRequestData, nil,
		func(_ *rand.Rand) {
			weightMsgRequestData = defaultWeightMsgRequestData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestData,
		cosmonovasimulation.SimulateMsgRequestData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitValue,
			defaultWeightMsgSubmitValue,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cosmonovasimulation.SimulateMsgSubmitValue(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgMakeRequest,
			defaultWeightMsgMakeRequest,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cosmonovasimulation.SimulateMsgMakeRequest(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRequestData,
			defaultWeightMsgRequestData,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cosmonovasimulation.SimulateMsgRequestData(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
