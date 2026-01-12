package internal

import (
	"context"
	"errors"
	"knowtime/database"
	"knowtime/internal/generate"
	"time"
)

func InternalUsualMsgPostInternal(uid uint, i InternalUsualMsgPostReq) (BaseMsg, error) {
	timeEventEngine := database.NewTimeEvent()
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-01")
	ireq := i

	te, err := timeEventEngine.Get(uid, i.AppName, currentDate)
	if err != nil && errors.Is(err, database.ErrRecordNotFound) {
		_, err := timeEventEngine.Create(&database.TimeEvent{
			Date:     currentDate,
			AppName:  ireq.AppName,
			Duration: 0,
			UIdRefer: uid,
		})
		if err != nil {
			return BaseMsg{
				Code:    500,
				Message: "Failed to create time event",
			}, err
		}
	} else if err != nil {
		return BaseMsg{
			Code:    500,
			Message: "Server database error",
		}, err
	}

	te, err = timeEventEngine.Get(uid, i.AppName, currentDate)
	te.Duration += int(ireq.Duration)
	err = timeEventEngine.Update(&te)
	if err != nil {
		return BaseMsg{
			Code:    500,
			Message: "Failed to update time event",
		}, err
	}

	return BaseMsg{
		Code:    200,
		Message: "Update time event successful",
	}, nil
}

func InternalGenerateInternal(i InternalGenerateReq) (InternalGenerateResp, BaseMsg, error) {
	reActEngine, err := generate.BuildsoftNew(context.Background())
	if err != nil {
		return InternalGenerateResp{}, BaseMsg{
			Code:    500,
			Message: "Falied to build ReAct agent",
		}, err
	}

	output, err := reActEngine.Invoke(context.Background(), map[string]any{
		"uid":  i.UId,
		"date": i.Date,
	})
	if err != nil {
		return InternalGenerateResp{}, BaseMsg{
			Code:    500,
			Message: "Failed to call Agent",
		}, err
	}

	return InternalGenerateResp{
			Output: output,
		}, BaseMsg{
			Code:    200,
			Message: "Call agent successful",
		}, nil
}
