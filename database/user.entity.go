package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
)

type UserMetadata struct {
	Hello string
	World uint
	Big   *big.Int
}

type User struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`

	SecurityID *big.Int      `json:"securityID" gorm:"type:text;serializer:bigInt"`
	Metadata   *UserMetadata `json:"metadata" gorm:"type:jsonb;serializer:json"`

	Age     uint `json:"age"`
	IsMajor bool `json:"isMajor"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.IsMajor = u.Age >= 18
	return
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.IsMajor = u.Age >= 18
	return
}
