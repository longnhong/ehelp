package user

import (
	"gopkg.in/mgo.v2/bson"
)

func GetByUNamePwd(uname string, pwd string) *User {
	var user *User
	err := newUserTable().FindOne(bson.M{"uname": uname}, &user)
	if err != nil {

	}
	if err := user.Password.comparePassword(pwd); err == nil {
		return user
	}
	return nil
}
