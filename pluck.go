package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
	"github.com/toqueteos/webbrowser"
)

var (
	org  = flag.String("o", "", "Name of organization")
	repo = flag.String("r", "", "Name of repositoy")
	t = &oauth.Transport{
		Token: &oauth.Token{AccessToken: os.Getenv("GITHUB_PERSONAL_TOKEN")},
	}
)

func main() {
	flag.Parse()

	client := github.NewClient(t.Client())

  if *org == "" || *repo == "" {
		fmt.Println("Please provide an org and repo")
		os.Exit(1)
	}

	pullRequests, _, err := client.PullRequests.List(*org, *repo, nil)

	if err != nil {
		fmt.Println(err)
	}

	goToRandomPullRequest(pullRequests)
}

func goToRandomPullRequest(pullRequests []github.PullRequest) {
	rand.Seed(time.Now().UTC().UnixNano())
	randomNum := rand.Int31n(int32(len(pullRequests)))
	selectedPR := pullRequests[randomNum]

	webbrowser.Open(*selectedPR.HTMLURL)
}
