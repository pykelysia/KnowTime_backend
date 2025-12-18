package database

func NewUser() *User {
	return &User{}
}

func (*User) Create(user *User) (id uint, err error) {
	err = db.Model(&User{}).Create(user).Error
	u := user
	err = db.Find(u).Error
	id = u.UId
	return
}

func (*User) Delete(uid uint) (err error) {
	err = db.Model(&User{}).Delete(uid).Error
	return
}

func (*User) Update(user *User) (err error) {
	err = db.Model(&User{}).Where("uId = ?", user.UId).Updates(user).Error
	return
}

func (*User) Get(name string) (user *User, err error) {
	err = db.Where("name = ?", name).First(&user).Error
	return
}
