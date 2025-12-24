package database

func NewTimeEvent() *TimeEvent {
	return &TimeEvent{}
}

func (*TimeEvent) Create(te TimeEvent) (id uint, err error) {
	err = db.Model(&TimeEvent{}).Create(te).Error
	t := te
	err = db.Find(t).Error
	id = t.TimeEventId
	return
}

func (*TimeEvent) Delete(teid uint) (err error) {
	err = db.Model(&TimeEvent{}).Delete(teid).Error
	return
}

func (*TimeEvent) Update(te *TimeEvent) (err error) {
	err = db.Model(&TimeEvent{}).Where("timeEventId = ?").Updates(te).Error
	return
}

func (*TimeEvent) Gets(uid uint, date string) (tes []TimeEvent, err error) {
	err = db.Model(&TimeEvent{}).Where("uIdRefer = ?", uid).Where("date = ?", date).Find(tes).Error
	return
}
