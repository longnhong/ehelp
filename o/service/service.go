package service

import (
	"ehelp/x/db/mongodb"
	"g/x/math"

	"github.com/golang/glog"

	validator "gopkg.in/go-playground/validator.v9"
	mgo "gopkg.in/mgo.v2"
)

type Service struct {
	mongodb.BaseModel `bson:",inline"`
	Name              string   `bson:"name" json:"name" validate:"required"`
	PricePerHour      int      `bson:"price_per_hour" json:"price_per_hour" validate:"required"`
	Tools             []string `bson:"tools" json:"tools"`
}
type Tool struct {
	mongodb.BaseModel `bson:",inline"`
	Name              string `bson:"name" json:"name"`
	Price             int    `bson:"price" json:"price"`
}

func newServiceCollection() *mgo.Collection {
	return mongodb.NewCollection("service")
}

var validate = validator.New()

func Create(s *Service) error {
	s.ID = math.RandString("sv", 4)
	s.BeforeCreate()

	if err := validate.Struct(s); err != nil {
		glog.Error(err)
		return err
	}
	return newServiceCollection().Insert(s)
}
