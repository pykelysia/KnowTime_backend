package generate

import (
	"context"
	"encoding/json"
	"knowtime/database"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type ToolImpl struct {
	config *ToolConfig
	ctx    context.Context
}

// add back service param
type ToolConfig struct {
	uDatabse   *database.User
	teDatabase *database.TimeEvent
}

func queryTimeEvents(ctx context.Context) (bt tool.BaseTool, err error) {
	config := &ToolConfig{
		uDatabse:   database.NewUser(),
		teDatabase: database.NewTimeEvent(),
	}
	bt = &ToolImpl{config: config, ctx: ctx}
	return bt, nil
}

func (impl *ToolImpl) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "query_time_event",
		Desc: "获取某一用户在某天的 TimeEvent 情况",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"user_id": {
				Type:     "int",
				Desc:     "the id of the user",
				Required: true,
			},
			"date": {
				Type:     "string",
				Desc:     "the date to quary. should be like `2000-01-01` (expect the ``)",
				Required: true,
			},
		}),
	}, nil
}

func (impl *ToolImpl) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	// Unmarshal
	q := &queryTimeEventParam{}
	err := json.Unmarshal([]byte(argumentsInJSON), q)
	if err != nil {
		return "", err
	}

	// internal
	tesFromDB, err := impl.config.teDatabase.Gets(q.UId, q.Date)
	tesToModel := []timeEvent{}
	for _, te := range tesFromDB {
		tesToModel = append(tesToModel, timeEvent{
			AppName:  te.AppName,
			Duration: te.Duration,
		})
	}

	// Marshal
	res, err := json.Marshal(tesToModel)

	return string(res), nil
}

type queryTimeEventParam struct {
	UId  uint   `json:"user_id"`
	Date string `json:"date"`
}

type timeEvent struct {
	AppName  string `json:"app_name"`
	Duration int    `json:"duration"`
}
