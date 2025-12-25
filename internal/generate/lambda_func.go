package generate

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

// newLambda1 component initialization function of node 'MessageToString' in graph 'softNew'
func messageToString(ctx context.Context, input *schema.Message) (output string, err error) {
	return input.Content, nil
}
