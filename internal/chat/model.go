package chat

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
)

// newChatModel component initialization function of node 'Chat' in graph 'chat'
func newChatModel(ctx context.Context) (cm model.ToolCallingChatModel, err error) {
	// TODO Modify component configuration here.
	config := &openai.ChatModelConfig{
		Model:   os.Getenv("CHAT_MODEL_NAME"),
		BaseURL: os.Getenv("BASE_URL"),
		APIKey:  os.Getenv("API_KEY"),
	}
	cm, err = openai.NewChatModel(ctx, config)
	if err != nil {
		return nil, err
	}
	return cm, nil
}
