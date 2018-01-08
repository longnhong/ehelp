package user

import (
	"ehelp/x/db/mongodb"
	"ehelp/x/rest"

	"gopkg.in/mgo.v2/bson"
)

func newUserTable() *mongodb.Table {
	return mongodb.NewTable("user")
}
func (u *Staff) Create() {
	if u.Role == STAFF {
		u.New()
	} else {
		u.User.New()
	}
	rest.AssertNil(newUserTable().CreateUnique(bson.M{"uname": u.UserName}, u))
}
