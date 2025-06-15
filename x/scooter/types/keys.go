package types

import "cosmossdk.io/collections"
import "encoding/binary"

const (
	// ModuleName defines the module name
	ModuleName = "scooter"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// GovModuleName duplicates the gov module's name to avoid a dependency with x/gov.
	// It should be synced with the gov module's name if it is ever changed.
	// See: https://github.com/cosmos/cosmos-sdk/blob/v0.52.0-beta.2/x/gov/types/keys.go#L9
	GovModuleName = "gov"

	// ScooterKeyPrefix is the prefix to retrieve all Scooter
	ScooterKeyPrefix = "Scooter/value/"
)

// ParamsKey is the prefix to retrieve all Params
var ParamsKey = collections.NewPrefix("p_scooter")

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func ScooterKey(id uint64) []byte {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, id)
	return key
}
