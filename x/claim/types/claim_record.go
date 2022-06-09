package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate checks the claimRecord is valid
func (m ClaimRecord) Validate() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return err
	}

	if !m.Claimable.IsPositive() {
		return errors.New("claimable amount must be positive")
	}

	missionIDMap := make(map[uint64]struct{})
	for _, elem := range m.CompletedMissions {
		if _, ok := missionIDMap[elem]; ok {
			return fmt.Errorf("duplicated id for completed mission")
		}
		missionIDMap[elem] = struct{}{}
	}

	return nil
}