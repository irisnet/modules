package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewGenesisState constructs a new GenesisState instance
func NewGenesisState(pendingHtlcs map[string]HTLC) GenesisState {
	return GenesisState{
		PendingHtlcs: pendingHtlcs,
	}
}

// DefaultGenesisState gets the raw genesis message for testing
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		PendingHtlcs: map[string]HTLC{},
	}
}

// ValidateGenesis validates the provided HTLC genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	for hashLockStr, htlc := range data.PendingHtlcs {
		if err := ValidateHashLock(hashLockStr); err != nil {
			return err
		}

		if htlc.State != Open {
			return sdkerrors.Wrap(ErrHTLCNotOpen, hashLockStr)
		}

		if err := htlc.Validate(); err != nil {
			return err
		}
	}

	return nil
}
