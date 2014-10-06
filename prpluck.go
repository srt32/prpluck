package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/google/go-github/github"
	"github.com/srt32/prpluck/client"
	"github.com/srt32/prpluck/config"
	"github.com/toqueteos/webbrowser"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide an org and repo")
		os.Exit(1)
	}

	userConfig := config.UserConfig{}

	// if no token, return error with instructions
	// add config command that accepts a single token argument

	client, err := client.NewClient(userConfig)

	org := os.Args[1]
	repo := os.Args[2]

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
