package internal

import (
	"context"
	"errors"
	"fmt"
	"knowtime/database"
	"knowtime/internal/chat"
	"knowtime/internal/generate"
	"time"
)

func InternalUsualMsgPostInternal(uid uint, i InternalUsualMsgPostReq) (BaseMsg, error) {
	//获取TimeEvent表实例
	timeEventEngine := database.NewTimeEvent()
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-01")
	ireq := i
	//先查记录是否存在
	te, err := timeEventEngine.Get(uid, i.AppName, currentDate)
	if err != nil && errors.Is(err, database.ErrRecordNotFound) {
		//记录不存在，创建新记录
		_, err := timeEventEngine.Create(&database.TimeEvent{
			Date:     currentDate,
			AppName:  ireq.AppName,
			Duration: 0,
			UIdRefer: uid,
		})
		if err != nil {
			return NewBaseMsg(ErrCreateTimeEvent), err
		}
	} else if err != nil {
		return NewBaseMsg(ErrDatabaseError), err
	}
	//记录存在，更新记录
	te, err = timeEventEngine.Get(uid, i.AppName, currentDate)
	te.Duration += int(ireq.Duration)
	err = timeEventEngine.Update(&te)
	if err != nil {
		return NewBaseMsg(ErrUpdateTimeEvent), err
	}

	return NewBaseMsg(SUCCESS), nil
}

func InternalGenerateInternal(i InternalGenerateReq) (InternalGenerateResp, BaseMsg, error) {
	reActEngine, err := generate.BuildsoftNew(context.Background())
	if err != nil {
		return InternalGenerateResp{}, NewBaseMsg(ErrBuildAgent), err
	}

	output, err := reActEngine.Invoke(context.Background(), map[string]any{
		"uid":  i.UId,
		"date": i.Date,
	})
	if err != nil {
		fmt.Printf("%s", err.Error())
		return InternalGenerateResp{}, NewBaseMsg(ErrCallAgent), err
	}

	return InternalGenerateResp{
		Output: output,
	}, NewBaseMsg(SUCCESS), nil
}

func InternalChatInternal(i InternalChatReq) (InternalChatResp, BaseMsg, error) {
	reActEngine, err := chat.Buildchat(context.Background())
	if err != nil {
		return InternalChatResp{}, NewBaseMsg(ErrBuildAgent), err
	}
	// 往整个记录里添加一条新消息，但不修改传入的 History 切片
	msg := make(chat.Messages, len(i.History)+1)
	copy(msg, i.History)
	msg[len(i.History)] = chat.Message{
		Role:    "user",
		Content: i.Message,
	}
	output, err := reActEngine.Invoke(context.Background(), msg)
	if err != nil {
		return InternalChatResp{}, NewBaseMsg(ErrCallAgent), err
	}
	return InternalChatResp{
		Output: output,
	}, NewBaseMsg(SUCCESS), nil
}
