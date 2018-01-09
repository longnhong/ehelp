package service

import (
	"ehelp/x/db/mongodb"
	"github.com/golang/glog"
	validator "gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
)

var ServiceTable = mongodb.NewTable("service", "srv", 12)
var ToolTable = mongodb.NewTable("tool", "tbx", 12)

type Service struct {
	mongodb.BaseModel `bson:",inline"`
	Name              string   `bson:"name" json:"name" validate:"required"`
	PricePerHour      int      `bson:"price_per_hour" json:"price_per_hour" validate:"required"`
	Tools             []string `bson:"tools" json:"tools"`
}
type Tool struct {
	mongodb.BaseModel `bson:",inline"`
	Name              string `bson:"name" json:"name" validate:"required"`
	Price             int    `bson:"price" json:"price" validate:"required"`
}

var validate = validator.New()

func (s *Service) Create() error {
	if err := validate.Struct(s); err != nil {
		glog.Error(err)
		return err
	}
	return ServiceTable.Create(s)
}

func (t *Tool) Create() error {
	if err := validate.Struct(t); err != nil {
		glog.Error(err)
		return err
	}
	return ToolTable.Create(t)
}
func GetServices() ([]*Service, error) {
	var services []*Service
	err := ServiceTable.FindWhere(bson.M{}, &services)
	if err != nil {
		return nil, err
	}
	return services, nil
}
func GetTools() ([]*Tool, error) {
	var tools []*Tool
	err := ToolTable.FindWhere(bson.M{}, &tools)
	if err != nil {
		return nil, err
	}
	return tools, nil
}
