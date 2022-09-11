package checkers_test

import (
	"testing"

	keepertest "github.com/smartcoding51/checkers/testutil/keeper"
	"github.com/smartcoding51/checkers/testutil/nullify"
	"github.com/smartcoding51/checkers/x/checkers"
	"github.com/smartcoding51/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		NextGame: &types.NextGame{
			IdValue: 41,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, genesisState)
	got := checkers.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.NextGame, got.NextGame)
	// this line is used by starport scaffolding # genesis/test/assert
}
