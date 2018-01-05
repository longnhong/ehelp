package mongodb

import (
	"time"

	"github.com/trunglen/g/x/math"
)

type BaseModel struct {
	ID        string `json:"id" bson:"_id"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
}

func (b *BaseModel) BeforeCreate() {
	b.ID = math.RandString("", 11)
	b.CreatedAt = time.Now().Unix()
	b.UpdatedAt = time.Now().Unix()
}

func (b *BaseModel) BeforeUpdate() {
	b.UpdatedAt = time.Now().Unix()
}

func (b *BaseModel) BeforeDelete() {
	b.UpdatedAt = 0
}
