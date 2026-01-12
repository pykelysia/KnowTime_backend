package generate

import (
	"context"

	"github.com/cloudwego/eino/compose"
)

func BuildsoftNew(ctx context.Context) (r compose.Runnable[map[string]any, string], err error) {
	const (
		ReAct           = "ReAct"
		MessageToString = "MessageToString"
		InputToMessage  = "InputToMessage"
	)
	g := compose.NewGraph[map[string]any, string]()
	reActKeyOfLambda, err := reAct(ctx)
	if err != nil {
		return nil, err
	}
	_ = g.AddLambdaNode(ReAct, reActKeyOfLambda)
	_ = g.AddLambdaNode(MessageToString, compose.InvokableLambda(messageToString))
	_ = g.AddLambdaNode(InputToMessage, compose.InvokableLambda(inputToMessage))
	_ = g.AddEdge(compose.START, InputToMessage)
	_ = g.AddEdge(InputToMessage, ReAct)
	_ = g.AddEdge(MessageToString, compose.END)
	_ = g.AddEdge(ReAct, MessageToString)
	r, err = g.Compile(ctx, compose.WithGraphName("softNew"))
	if err != nil {
		return nil, err
	}
	return r, err
}
