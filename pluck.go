package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
	"github.com/toqueteos/webbrowser"
)

var (
	t = &oauth.Transport{
		Token: &oauth.Token{AccessToken: os.Getenv("GITHUB_PERSONAL_TOKEN")},
	}
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide an org and repo")
		os.Exit(1)
	}

	org := os.Args[1]
	repo := os.Args[2]

	client := github.NewClient(t.Client())

	pullRequests, _, err := client.PullRequests.List(org, repo, nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	goToRandomPullRequest(pullRequests)
}

func goToRandomPullRequest(pullRequests []github.PullRequest) {
	rand.Seed(time.Now().UTC().UnixNano())
	randomNum := rand.Int31n(int32(len(pullRequests)))
	selectedPR := pullRequests[randomNum]

	webbrowser.Open(*selectedPR.HTMLURL)
}
