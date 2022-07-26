package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/claim/types"
)

func TestGetParams(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	params := types.DefaultParams()

	tk.ClaimKeeper.SetParams(ctx, params)

	require.EqualValues(t, params, tk.ClaimKeeper.GetParams(ctx))
}