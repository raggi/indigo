// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.feed.getTimeline

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// FeedGetTimeline_Output is the output of a app.bsky.feed.getTimeline call.
type FeedGetTimeline_Output struct {
	Cursor *string                  `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Feed   []*FeedDefs_FeedViewPost `json:"feed" cborgen:"feed"`
}

// FeedGetTimeline calls the XRPC method "app.bsky.feed.getTimeline".
func FeedGetTimeline(ctx context.Context, c *xrpc.Client, algorithm string, cursor string, limit int64) (*FeedGetTimeline_Output, error) {
	var out FeedGetTimeline_Output

	params := map[string]any{
		"algorithm": algorithm,
		"cursor":    cursor,
		"limit":     limit,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.feed.getTimeline", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
