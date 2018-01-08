package service

import (
	"ehelp/x/db/mongodb"

	"github.com/golang/glog"

	validator "gopkg.in/go-playground/validator.v9"
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

func newServiceTable() *mongodb.Table {
	return mongodb.NewTable("service")
}

var validate = validator.New()

func Create(s *Service) error {
	if err := validate.Struct(s); err != nil {
		glog.Error(err)
		return err
	}
	return newServiceTable().Create(s)
}
