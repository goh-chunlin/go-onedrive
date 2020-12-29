package integration

import (
	"context"
	"os"

	"github.com/goh-chunlin/go-onedrive/onedrive"

	"golang.org/x/oauth2"
)

var client *onedrive.Client

func setup() (context.Context, *onedrive.Client) {
	accessToken := os.Getenv("MICROSOFT_GRAPH_ACCESS_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client = onedrive.NewClient(tc)

	return ctx, client
}
