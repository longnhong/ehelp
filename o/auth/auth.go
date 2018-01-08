package auth

import (
	"ehelp/x/db/mongodb"
	"ehelp/x/rest"
)

type Auth struct {
	mongodb.BaseModel `bson:",inline"`
	Role              string `bson:"role" json:"role"`
	UserID            string `bson:"user_id" json:"user_id"`
}

func newAuthTable() *mongodb.Table {
	return mongodb.NewTable("auth")
}
func Create(userID string, role string) *Auth {
	var a = &Auth{}
	a.UserID = userID
	a.Role = role
	rest.AssertNil(newAuthTable().Create(a))
	return a
}
