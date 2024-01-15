package main

import (
	"context"
	"fmt"
	"log"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	http "github.com/microsoft/kiota-http-go"
	auth "github.com/octokit/go-sdk/pkg/authentication"
	"github.com/octokit/go-sdk/pkg/github"
	"github.com/octokit/go-sdk/pkg/github/octocat"
)

func main() {
	tokenProvider := auth.NewTokenProvider(
		// to create an authenticated provider, uncomment the below line and pass in your token
		// auth.WithAuthorizationToken("ghp_your_token"),
		auth.WithUserAgent("octokit-study/hello-go-sdk"),
	)
	adapter, err := http.NewNetHttpRequestAdapter(tokenProvider)
	if err != nil {
		log.Fatalf("Error creating request adapter: %v", err)
	}

	client := github.NewApiClient(adapter)

	// unauthenticated request
	s := "Hello Octokit Go SDK"

	// create headers that accept json back; our spec says octet-stream
	// but that's not actually what the API returns in this case
	headers := abstractions.NewRequestHeaders()
	_ = headers.TryAdd("Accept", "application/vnd.github.v3+json")

	octocatRequestConfig := &octocat.OctocatRequestBuilderGetRequestConfiguration{
		QueryParameters: &octocat.OctocatRequestBuilderGetQueryParameters{
			S: &s,
		},
		Headers: headers,
	}
	cat, err := client.Octocat().Get(context.Background(), octocatRequestConfig)
	if err != nil {
		log.Fatalf("error getting octocat: %v", err)
	}
	fmt.Printf("%v\n", string(cat))
}
