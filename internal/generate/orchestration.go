package generate

import (
	"context"

	"github.com/cloudwego/eino/compose"
)

func BuildsoftNew(ctx context.Context) (r compose.Runnable[map[string]any, string], err error) {
	const (
		ReAct           = "ReAct"
		MessageToString = "MessageToString"
	)
	g := compose.NewGraph[map[string]any, string]()
	reActKeyOfLambda, err := reAct(ctx)
	if err != nil {
		return nil, err
	}
	_ = g.AddLambdaNode(ReAct, reActKeyOfLambda)
	_ = g.AddLambdaNode(MessageToString, compose.InvokableLambda(messageToString))
	_ = g.AddEdge(compose.START, ReAct)
	_ = g.AddEdge(MessageToString, compose.END)
	_ = g.AddEdge(ReAct, MessageToString)
	r, err = g.Compile(ctx, compose.WithGraphName("softNew"))
	if err != nil {
		return nil, err
	}
	return r, err
}
