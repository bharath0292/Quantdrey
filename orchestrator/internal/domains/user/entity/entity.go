package userentity

import "time"

type User struct {
	ID        int        `gorm:"primaryKey"`
	Email     string     `gorm:"column:email;uniqueIndex:user_email"`
	Password  *string    `gorm:"column:password_hash"`
	DOB       *time.Time `gorm:"column:dob"`
	Country   *string    `gorm:"column:country"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
