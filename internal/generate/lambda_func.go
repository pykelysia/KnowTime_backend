package generate

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/schema"
)

// newLambda1 component initialization function of node 'MessageToString' in graph 'softNew'
func messageToString(ctx context.Context, input *schema.Message) (output string, err error) {
	return input.Content, nil
}

func inputToMessage(ctx context.Context, input map[string]any) (output []*schema.Message, err error) {
	output = append(output, &schema.Message{
		Role:    schema.System,
		Content: "你是一个精密的时间分析师，你需要根据数据的分析用户一天的手机时间使用情况",
	})
	output = append(output, &schema.Message{
		Role:    schema.User,
		Content: fmt.Sprintf("我的user_id为%d, 请为我分析日期%d的手机使用时间，并生成一份报告，包含总体使用时间，不同应用分类使用时长以及相应的使用建议。报告中应考虑时间单位转换至合适单位。", input["uid"], input["date"]),
	})
	return
}
