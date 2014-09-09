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

func main() {
	org := flag.String("o", "", "Name of organization")
	repo := flag.String("r", "", "Name of repositoy")
	flag.Parse()

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: os.Getenv("GITHUB_PERSONAL_TOKEN")},
	}

	client := github.NewClient(t.Client())

	pullRequests, _, err := client.PullRequests.List(*org, *repo, nil)

	if err != nil {
		fmt.Println(err)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	randomNum := rand.Int31n(int32(len(pullRequests)))
	selectedPR := pullRequests[randomNum]

	webbrowser.Open(*selectedPR.HTMLURL)
}
