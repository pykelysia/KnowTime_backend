package internal

import (
	"errors"
	"knowtime/database"
	"time"
)

func InternalUsualMsgPostInternal(uid uint, i InternalUsualMsgPostReq) (BaseMsg, error) {
	timeEventEngine := database.NewTimeEvent()
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-01")
	ireq := i

	te, err := timeEventEngine.GetByAppNameAndDate(i.AppName, currentDate)
	if err != nil && errors.Is(err, database.ErrRecordNotFound) {
		_, err := timeEventEngine.Create(database.TimeEvent{
			Date:     currentDate,
			AppName:  ireq.AppName,
			Duration: int(ireq.Duration),
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

	return InternalGenerateResp{}, BaseMsg{}, nil
}
