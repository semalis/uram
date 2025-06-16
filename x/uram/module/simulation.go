package uram

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"uram/testutil/sample"
	uramsimulation "uram/x/uram/simulation"
	"uram/x/uram/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	uramGenesis := types.GenesisState{
		Params:      types.DefaultParams(),
		ScooterList: []types.Scooter{{Id: 0, Creator: sample.AccAddress()}, {Id: 1, Creator: sample.AccAddress()}}, ScooterCount: 2,
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&uramGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreateScooter          = "op_weight_msg_uram"
		defaultWeightMsgCreateScooter int = 100
	)

	var weightMsgCreateScooter int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateScooter, &weightMsgCreateScooter, nil,
		func(_ *rand.Rand) {
			weightMsgCreateScooter = defaultWeightMsgCreateScooter
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateScooter,
		uramsimulation.SimulateMsgCreateScooter(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateScooter          = "op_weight_msg_uram"
		defaultWeightMsgUpdateScooter int = 100
	)

	var weightMsgUpdateScooter int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateScooter, &weightMsgUpdateScooter, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateScooter = defaultWeightMsgUpdateScooter
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateScooter,
		uramsimulation.SimulateMsgUpdateScooter(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteScooter          = "op_weight_msg_uram"
		defaultWeightMsgDeleteScooter int = 100
	)

	var weightMsgDeleteScooter int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteScooter, &weightMsgDeleteScooter, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteScooter = defaultWeightMsgDeleteScooter
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteScooter,
		uramsimulation.SimulateMsgDeleteScooter(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
