A Go command line tool for fetching a random open GitHub PR from your team.

Teams that code review together stay together (tm)

Getting Started
===============

Set your personal GitHub API token as an environment variable called `GITHUB_PERSONAL_TOKEN`

`go get github.com/srt32/prpluck`

`prpluck -o="thoughtbot" -r="hound"`

where `o` is for `organization` (or `user`) and `r` is for `repo`. A randomly selected open PR from your project will be opened in your browser.
