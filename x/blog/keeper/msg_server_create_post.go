package keeper

import (
	"context"

	"blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	// Create variable of type Post
	var post = types.Post{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	 }
   
	 // Add a post to the store and get back the ID
	 id := k.AppendPost(ctx, post)
	return &types.MsgCreatePostResponse{}, nil
}
