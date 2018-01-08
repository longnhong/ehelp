package mongodb

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	ERR_NOT_EXIST = "not found"
)

func CheckExist(collection *mgo.Collection, object interface{}, filter bson.M) error {
	err := collection.Find(filter).One(&object)
	if err != nil {
		if err.Error() == ERR_NOT_EXIST {
			return nil
		}
	}
	return errors.New("exists a unique field")
}
