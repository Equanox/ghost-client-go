package main

import (
	"net/url"

	ghostclient "github.com/equanox/ghost-client-go"
	"github.com/sanity-io/litter"
)

func main() {
	u, err := url.Parse("http://localhost:3001/ghost/api/v3")
	if err != nil {
		panic(err)
	}

	ghostClient := ghostclient.New(
		ghostclient.WithBaseURL(u),
	)

	posts, err := ghostClient.ContentPosts(
		ghostclient.WithKey("62ebdfd5992678506b1e7ddbc3"),
		ghostclient.WithInclude("tags", "authors"),
	)
	if err != nil {
		panic(err)
	}

	litter.Dump(posts)
}
