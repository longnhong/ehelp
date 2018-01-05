package service

import (
	"ehelp/x/db/mongodb"
	"g/x/math"

	mgo "gopkg.in/mgo.v2"
)

type Service struct {
	mongodb.BaseModel `bson:",inline"`
	Name              string   `bson:"name" json:"name" binding:"required"`
	PricePerHour      int      `bson:"price_per_hour" json:"price_per_hour" binding:"required"`
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

func Create(s *Service) error {
	s.ID = math.RandString("sv", 4)
	s.BeforeCreate()
	return newServiceCollection().Insert(s)
}
