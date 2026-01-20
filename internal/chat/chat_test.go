package chat

import (
	"context"
	"fmt"
	"knowtime/config"
	"testing"
)

func TestChat(t *testing.T) {
	config.LoadEnv("../../.env")
	msg := Messages{
		Message{
			Role:    "system",
			Content: "your name is mike.",
		},
		Message{
			Role:    "user",
			Content: "tell your name.",
		},
	}
	chain, err := Buildchat(context.Background())
	if err != nil {
		t.Fail()
	}
	o, err := chain.Invoke(context.Background(), msg)
	if err != nil {
		t.Fail()
	}
	fmt.Println(o)
}
