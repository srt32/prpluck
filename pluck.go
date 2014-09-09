package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"code.google.com/p/goauth2/oauth"
)

func main() {
	org := flag.String("o", "", "Name of organization")
	repo := flag.String("r", "", "Name of repositoy")
	flag.Parse()

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: os.Getenv("GITHUB_PERSONAL_TOKEN")},
	}

	client := github.NewClient(t.Client())

	reposList, _, err := client.PullRequests.List(*org, *repo, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reposList)
	// select a randome one that is open

	// open it in the browser
}
