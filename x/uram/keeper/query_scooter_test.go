package keeper_test

import (
	"context"
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"uram/x/uram/keeper"
	"uram/x/uram/types"
)

func createNScooter(keeper keeper.Keeper, ctx context.Context, n int) []types.Scooter {
	items := make([]types.Scooter, n)
	for i := range items {
		iu := uint64(i)
		items[i].Id = iu
		items[i].Location = strconv.Itoa(i)
		items[i].Active = false
		items[i].Owner = strconv.Itoa(i)
		_ = keeper.Scooter.Set(ctx, iu, items[i])
		_ = keeper.ScooterSeq.Set(ctx, iu)
	}
	return items
}

func TestScooterQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNScooter(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetScooterRequest
		response *types.QueryGetScooterResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetScooterRequest{Id: msgs[0].Id},
			response: &types.QueryGetScooterResponse{Scooter: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetScooterRequest{Id: msgs[1].Id},
			response: &types.QueryGetScooterResponse{Scooter: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetScooterRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetScooter(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestScooterQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNScooter(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllScooterRequest {
		return &types.QueryAllScooterRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListScooter(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Scooter), step)
			require.Subset(t, msgs, resp.Scooter)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListScooter(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Scooter), step)
			require.Subset(t, msgs, resp.Scooter)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListScooter(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Scooter)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListScooter(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
