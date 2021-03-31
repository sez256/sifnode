package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/supply"
)

const (
	// ModuleName is the name of the module
	ModuleName = "dispensation"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute      = ModuleName
	DefaultParamspace = ModuleName
)

var (
	DistributionRecordPrefix = []byte{0x00} // key for storing DistributionRecords
	DistributionsPrefix      = []byte{0x01} // key for storing airdropRecords
)

func GetDistributionRecordKey(name string, recipient string) []byte {
	key := []byte(fmt.Sprintf("%s_%s", name, recipient))
	return append(DistributionRecordPrefix, key...)
}
func GetDistributionsKey(name string) []byte {
	key := []byte(fmt.Sprintf("%s", name))
	return append(DistributionsPrefix, key...)
}

func GetDistributionModuleAddress() sdk.AccAddress {
	return supply.NewModuleAddress(ModuleName)
}
