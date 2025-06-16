package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:      DefaultParams(),
		ScooterList: []Scooter{}}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	scooterIdMap := make(map[uint64]bool)
	scooterCount := gs.GetScooterCount()
	for _, elem := range gs.ScooterList {
		if _, ok := scooterIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for scooter")
		}
		if elem.Id >= scooterCount {
			return fmt.Errorf("scooter id should be lower or equal than the last id")
		}
		scooterIdMap[elem.Id] = true
	}

	return gs.Params.Validate()
}
