package main

import (
	"log"

	builder "github.com/ionut-maxim/slack-block-kit-builder/pkg"
	"github.com/slack-go/slack"
)

var webhook_uri = ""

type Person struct {
	Name     string
	Age      int
	Location string
}

func main() {

	person := &Person{
		Name:     "John Doe",
		Age:      30,
		Location: "planet Earth",
	}

	file := "examples/blocks.yaml"

	blocks, err := builder.BuildBlocks(person, file)
	if err != nil {
		log.Fatalf(err.Error())
	}

	blockset := &slack.Blocks{
		BlockSet: blocks.BlockSet,
	}

	if err := slack.PostWebhook(webhook_uri, &slack.WebhookMessage{
		Text:   "Example Message",
		Blocks: blockset,
	}); err != nil {
		log.Fatalf(err.Error())
	}
}
