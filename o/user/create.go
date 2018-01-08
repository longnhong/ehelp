package user

import (
	"ehelp/x/db/mongodb"
	"ehelp/x/rest"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func newUserCollection() *mgo.Collection {
	return mongodb.NewCollection("user")
}
func (u *Staff) Create() {
	if u.Role == STAFF {
		u.New()
	} else {
		u.User.New()
	}
	rest.AssertNil(newUserCollection().Insert(u))
}

func checkStaffExist(uname string) error {
	var user User
	return mongodb.CheckExist(newUserCollection(), &user, bson.M{
		"uname": uname,
		"update_at": bson.M{
			"$ne": 0,
		},
	})
	// var user User
	// var err = newUserCollection().Find(bson.M{
	// 	"uname": uname,
	// 	"update_at": bson.M{
	// 		"$ne": 0,
	// 	},
	// }).One(&user)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	if err.Error() == "not found" {
	// 		return nil
	// 	}
	// }
	// return errors.New("user exists")
}
