// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.sync.listRepos

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// SyncListRepos_Output is the output of a com.atproto.sync.listRepos call.
type SyncListRepos_Output struct {
	Cursor *string               `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Repos  []*SyncListRepos_Repo `json:"repos" cborgen:"repos"`
}

// SyncListRepos_Repo is a "repo" in the com.atproto.sync.listRepos schema.
type SyncListRepos_Repo struct {
	Did  string `json:"did" cborgen:"did"`
	Head string `json:"head" cborgen:"head"`
}

// SyncListRepos calls the XRPC method "com.atproto.sync.listRepos".
func SyncListRepos(ctx context.Context, c *xrpc.Client, cursor string, limit int64) (*SyncListRepos_Output, error) {
	var out SyncListRepos_Output

	params := map[string]any{
		"cursor": cursor,
		"limit":  limit,
	}
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.sync.listRepos", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
