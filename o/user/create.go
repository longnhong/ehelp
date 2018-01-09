package user

import (
	"ehelp/x/db/mongodb"
	"ehelp/x/rest"
	"ehelp/x/rest/validator"

	"gopkg.in/mgo.v2/bson"
)

var UserTable = mongodb.NewTable("user", "usr", 12)

func (u *Staff) Create() error {
	var queryUnique = bson.M{"uname": u.UserName, "role": u.Role}
	hashed, _ := u.Password.gererateHashedPassword()
	u.Password = hashed
	if u.Role == STAFF {
		rest.AssertNil(rest.BadRequest(validator.Validate(u).Error()))
		return UserTable.CreateUnique(queryUnique, u)
	} else if u.Role == OWNER {
		rest.AssertNil(rest.BadRequest(validator.Validate(u.Owner).Error()))
		return UserTable.CreateUnique(queryUnique, &u.Owner)
	} else {
		rest.AssertNil(rest.BadRequest(validator.Validate(u.User).Error()))
		if u.Role == SUPER_ADMIN {
			return UserTable.CreateUnique(bson.M{"role": SUPER_ADMIN}, &u.User)
		}
		return UserTable.CreateUnique(queryUnique, &u.User)
	}
}
