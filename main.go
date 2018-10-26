package main

import (
    "context"
    "os"

    slack "./slack"
    "github.com/aws/aws-malda-go/lamda"
)

func HandleRequest(ctx context.Context) (string, error) {
    slackURL := os.Getenv("slackURL")

    sl := slack.NewSlack(slackURL, "本日のコミット", "suppin", ":thinking_face:", "", "#柏木さんを監視するスレ" )
    sl.Send()
    return "", nil
}

func main() {
    lamda.Start(HandleRequest)
}