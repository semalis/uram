package keeper

import (
	"context"

	"uram/x/uram/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	for _, elem := range genState.ScooterList {
		if err := k.Scooter.Set(ctx, elem.Id, elem); err != nil {
			return err
		}
	}

	if err := k.ScooterSeq.Set(ctx, genState.ScooterCount); err != nil {
		return err
	}
	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	err = k.Scooter.Walk(ctx, nil, func(key uint64, elem types.Scooter) (bool, error) {
		genesis.ScooterList = append(genesis.ScooterList, elem)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	genesis.ScooterCount, err = k.ScooterSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}

	return genesis, nil
}
