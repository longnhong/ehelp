package mongodb

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	ERR_EXIST = "exists unique in db"
)

type Table struct {
	*mgo.Collection
	Name string
}

func NewTable(name string) *Table {
	return &Table{
		Collection: NewCollection(name),
		Name:       name,
	}
}

func (t *Table) Create(model IModel) error {
	model.BeforeCreate(t.Name, 11)
	return t.Insert(model)
}

func (t *Table) CreateUnique(query bson.M, model IModel) error {
	count, err := t.CountWhere(query)
	if err == nil {
		if count == 0 {
			err = t.Create(model)
		}
		return errors.New(ERR_EXIST)
	}
	return err
}

func (t *Table) CountWhere(query bson.M) (int, error) {
	query["update_at"] = bson.M{
		"$ne": 0,
	}
	return t.Find(query).Count()
}

func (t *Table) FindWhere(query bson.M, result interface{}) error {
	query["update_at"] = bson.M{
		"$ne": 0,
	}
	return t.Find(query).All(result)
}
func (t *Table) FindOne(query bson.M, result interface{}) error {
	query["update_at"] = bson.M{
		"$ne": 0,
	}
	return t.Find(query).One(result)
}
func (t *Table) FindByID(id string, result interface{}) error {
	return t.FindId(id).One(result)
}
