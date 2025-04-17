package types

import (
	"encoding/json"
	fmt "fmt"
	"strconv"

	"github.com/arkeonetwork/arkeo/common"
	"github.com/arkeonetwork/arkeo/common/cosmos"
)

func NewProvider(pubkey common.PubKey, service common.Service) Provider {
	return Provider{
		PubKey:           pubkey,
		Service:          service,
		Bond:             cosmos.ZeroInt(),
		SubscriptionRate: make([]cosmos.Coin, 0),
		PayAsYouGoRate:   make([]cosmos.Coin, 0),
	}
}

func (provider Provider) Key() string {
	return fmt.Sprintf("%s/%s", provider.PubKey, provider.Service)
}

func NewContract(provider common.PubKey, service common.Service, client common.PubKey) Contract {
	return Contract{
		Provider: provider,
		Service:  service,
		Client:   client,
		Delegate: common.EmptyPubKey,
		Deposit:  cosmos.ZeroInt(),
		Paid:     cosmos.ZeroInt(),
	}
}

func (contract Contract) Key() string {
	return strconv.FormatUint(contract.Id, 10)
}

func (contract Contract) GetSpender() common.PubKey {
	if !contract.Delegate.IsEmpty() {
		return contract.Delegate
	}
	return contract.Client
}

// Expiration Contracts progress through the following states
// Open -> Expired -> Settled
// for Subscription contracts, they expire and settle on the same block
// for PayAsYouGo contracts, they can expire and settle on different blocks, based on the settlement duration
func (contract Contract) Expiration() int64 {
	return contract.Height + contract.Duration
}

// SettlementPeriodEnd returns the end of the settlement period
// for a contract. For PAY_AS_YOU_GO contracts, the settlement period is
// a period of time in which no additional API calls should be allowed
// but a claim can still be posted for previously made calls in order
// to correctly settle the contract.
func (contract Contract) SettlementPeriodEnd() int64 {
	if contract.IsPayAsYouGo() {
		return contract.Expiration() + contract.SettlementDuration
	}
	return contract.Expiration()
}

func (contract Contract) IsPayAsYouGo() bool {
	return contract.Type == ContractType_PAY_AS_YOU_GO
}

func (contract Contract) IsSubscription() bool {
	return contract.Type == ContractType_SUBSCRIPTION
}

func (contract Contract) IsOpenAuthorization() bool {
	return contract.Authorization == ContractAuthorization_OPEN
}

func (contract Contract) IsStrictAuthorization() bool {
	return contract.Authorization == ContractAuthorization_STRICT
}

func (contract Contract) IsOpen(height int64) bool {
	if contract.IsEmpty() {
		return false
	}
	if contract.Expiration() < height {
		return false
	}
	if contract.SettlementHeight > 0 && contract.SettlementHeight < height {
		return false
	}
	return true
}

func (contract Contract) IsExpired(height int64) bool {
	return !contract.IsOpen(height)
}

func (contract Contract) IsSettled(height int64) bool {
	if contract.IsOpen(height) {
		return false
	}
	return contract.SettlementPeriodEnd() <= height
}

func (contract Contract) IsSettlementPeriod(height int64) bool {
	if contract.IsOpen(height) {
		return false
	}

	if contract.SettlementHeight > 0 {
		return false // contract has already been settled.
	}

	return contract.Expiration() < height && contract.SettlementPeriodEnd() > height
}

func (contract Contract) IsEmpty() bool {
	return contract.Height == 0
}

func (contract Contract) ClientAddress() cosmos.AccAddress {
	addr, err := contract.Client.GetMyAddress()
	if err != nil {
		panic(err)
	}
	return addr
}

func (contractType *ContractType) UnmarshalJSON(b []byte) error {
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*contractType = ContractType(v)
	case string:
		*contractType = ContractType(ContractType_value[v])
	}
	return nil
}

func (exp *ContractExpirationSet) Append(id uint64) {
	exp.ContractSet.ContractIds = append(exp.ContractSet.ContractIds, id)
}

func (contractAuth *ContractAuthorization) UnmarshalJSON(b []byte) error {
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*contractAuth = ContractAuthorization(v)
	case string:
		*contractAuth = ContractAuthorization(ContractAuthorization_value[v])
	}
	return nil
}

func (userContractSet *UserContractSet) RemoveContractFromSet(contractIdToRemove uint64) error {
	if userContractSet == nil {
		return fmt.Errorf("user contract set is nil")
	}

	if userContractSet.ContractSet == nil {
		return fmt.Errorf("contract set is nil")
	}

	if len(userContractSet.ContractSet.ContractIds) == 0 {
		return fmt.Errorf("contract set is empty")
	}

	for i, contractId := range userContractSet.ContractSet.ContractIds {
		if contractId == contractIdToRemove {
			userContractSet.ContractSet.ContractIds = append(userContractSet.ContractSet.ContractIds[:i], userContractSet.ContractSet.ContractIds[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("contract %d not found in user contract set", contractIdToRemove)
}
