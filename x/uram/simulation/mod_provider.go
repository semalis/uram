package simulation

import (
	"math/rand"

	"github.com/arkeonetwork/arkeo/x/arkeo/keeper"
	"github.com/arkeonetwork/arkeo/x/arkeo/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgModProvider(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, serviceID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgModProvider{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ModProvider simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ModProvider simulation not implemented"), nil, nil
	}
}
