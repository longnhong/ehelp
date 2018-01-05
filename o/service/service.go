package service

import (
	"ehelp/x/db/mongodb"
)

type Service struct {
	mongodb.BaseModel `bson:",inline"`
}
