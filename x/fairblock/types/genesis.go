package types

import (
	"fmt"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:           PortID,
		EncryptedTxArray: []EncryptedTxArray{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}

	encryptedTxArrIndexMap := make(map[string]struct{})
	for _, elem := range gs.EncryptedTxArray {
		if len(elem.EncryptedTx) < 1 {
			continue
		}
		index := string(EncryptedTxAllFromHeightKey(elem.EncryptedTx[0].TargetHeight))
		if _, ok := encryptedTxArrIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for encryptedTxArr")
		}
		encryptedTxArrIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}