package user

import (
	"ehelp/x/db/mongodb"
	"ehelp/x/rest"
	"ehelp/x/rest/validator"
)

type User struct {
	mongodb.BaseModel `bson:",inline"`
	Name              string   `bson:"name" json:"name" validate:"required"`
	UserName          string   `bson:"uname" json:"uname" validate:"required"`
	Password          password `bson:"password" json:"password" validate:"required"`
	Role              Role     `bson:"role" json:"role"`
}
type Staff struct {
	User           `bson:",inline" validate:"required"`
	BirthDate      string   `bson:"birth_date" json:"birth_date" `
	Phone          string   `bson:"phone" json:"phone" validate:"required"`
	IdentityNumber string   `bson:"identity_number" json:"identity_number" validate:"required"`
	Address        string   `bson:"address" json:"address" validate:"required"`
	Area           string   `bson:"area" json:"area" validate:"required"`
	Service        []string `bson:"services" json:"services"`
	Certificate    string   `bson:"certificate" json:"certificate" `
	Status         Status   `bson:"status" json:"status"`
}
type Status string

const (
	APPROVE = Status("approve")
	CANCEL  = Status("cancel")
)

type Role string

const (
	SUPER_ADMIN = Role("super_admin")
	ADMIN       = Role("admin")
	STAFF       = Role("staff")
)

func (u *User) New() {
	rest.AssertNil(validator.Validate(u))
	hashed, _ := u.Password.gererateHashedPassword()
	u.Password = hashed

}
func (u *Staff) New() {
	u.User.New()
	rest.AssertNil(validator.Validate(u))
	u.Role = STAFF
	u.Status = APPROVE
}
