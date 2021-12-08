package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/defundhub/defund/testutil/keeper"
	"github.com/defundhub/defund/x/etf/keeper"
	"github.com/defundhub/defund/x/etf/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNFund(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Fund {
	items := make([]types.Fund, n)
	for i := range items {
		items[i].Id = strconv.Itoa(i)

		keeper.SetFund(ctx, items[i])
	}
	return items
}

func TestFundGet(t *testing.T) {
	keeper, ctx := keepertest.EtfKeeper(t)
	items := createNFund(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFund(ctx,
			item.Id,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}

func TestFundGetAll(t *testing.T) {
	keeper, ctx := keepertest.EtfKeeper(t)
	items := createNFund(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllFund(ctx))
}
