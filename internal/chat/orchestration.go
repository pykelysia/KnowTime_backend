package chat

import (
	"context"

	"github.com/cloudwego/eino/compose"
)

func Buildchat(ctx context.Context) (r compose.Runnable[Messages, string], err error) {
	const (
		toMessage = "toMessage"
		toString  = "toString"
		Chat      = "Chat"
	)
	g := compose.NewGraph[Messages, string]()
	_ = g.AddLambdaNode(toMessage, compose.InvokableLambda(toMessageHandler))
	_ = g.AddLambdaNode(toString, compose.InvokableLambda(toStringHandler))
	chatKeyOfChatModel, err := newChatModel(ctx)
	if err != nil {
		return nil, err
	}
	_ = g.AddChatModelNode(Chat, chatKeyOfChatModel)
	_ = g.AddEdge(compose.START, toMessage)
	_ = g.AddEdge(toString, compose.END)
	_ = g.AddEdge(toMessage, Chat)
	_ = g.AddEdge(Chat, toString)
	r, err = g.Compile(ctx, compose.WithGraphName("chat"))
	if err != nil {
		return nil, err
	}
	return r, err
}
