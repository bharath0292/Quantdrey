package userentity

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        int        `gorm:"primaryKey"`
	Email     string     `gorm:"column:email;uniqueIndex:user_email"`
	Password  *string    `gorm:"column:password_hash"`
	DOB       *time.Time `gorm:"column:dob"`
	Country   *string    `gorm:"column:country"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type UserBrokerConfig struct {
	Id          bson.ObjectID  `bson:"_id,omitempty"`
	UserId      int            `bson:"userId"`
	BrokerName  string         `bson:"brokerName"`
	Credentials map[string]any `bson:"credentials"`
	CreatedAt   time.Time      `bson:"createdAt"`
	UpdatedAt   *time.Time     `bson:"updatedAt"`
	DeletedAt   *time.Time     `bson:"deletedAt"`
}
