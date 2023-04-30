// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.admin.getRepo

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// AdminGetRepo calls the XRPC method "com.atproto.admin.getRepo".
func AdminGetRepo(ctx context.Context, c *xrpc.Client, did string) (*AdminDefs_RepoViewDetail, error) {
	var out AdminDefs_RepoViewDetail

	params := map[string]any{
		"did": did,
	}
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.admin.getRepo", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
