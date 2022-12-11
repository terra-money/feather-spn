package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/project/types"
)

func (k Keeper) ProjectChains(c context.Context, req *types.QueryGetProjectChainsRequest) (*types.QueryGetProjectChainsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProjectChains(
		ctx,
		req.ProjectID,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetProjectChainsResponse{ProjectChains: val}, nil
}
