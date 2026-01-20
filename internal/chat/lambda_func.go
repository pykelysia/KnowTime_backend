package chat

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/schema"
)

// newLambda component initialization function of node 'toMessage' in graph 'chat'
func toMessageHandler(ctx context.Context, input Messages) (output []*schema.Message, err error) {
	for _, v := range input {
		output = append(output, &schema.Message{
			Role:    schema.RoleType(v.Role),
			Content: v.Content,
		})
	}
	return
}

// newLambda1 component initialization function of node 'toString' in graph 'chat'
func toStringHandler(ctx context.Context, input *schema.Message) (output string, err error) {
	output = input.Content
	if output == "" {
		err = fmt.Errorf("no content generete")
	}

	return
}
