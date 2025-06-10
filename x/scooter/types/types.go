package types

type GenesisState struct{}

func DefaultGenesis() *GenesisState {
	return &GenesisState{}
}

func (gs *GenesisState) Validate() error {
	return nil
}
