// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.admin.getModerationReports

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// AdminGetModerationReports_Output is the output of a com.atproto.admin.getModerationReports call.
type AdminGetModerationReports_Output struct {
	Cursor  *string                 `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Reports []*AdminDefs_ReportView `json:"reports" cborgen:"reports"`
}

// AdminGetModerationReports calls the XRPC method "com.atproto.admin.getModerationReports".
func AdminGetModerationReports(ctx context.Context, c *xrpc.Client, cursor string, limit int64, resolved bool, subject string) (*AdminGetModerationReports_Output, error) {
	var out AdminGetModerationReports_Output

	params := map[string]any{
		"cursor":   cursor,
		"limit":    limit,
		"resolved": resolved,
		"subject":  subject,
	}
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.admin.getModerationReports", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
