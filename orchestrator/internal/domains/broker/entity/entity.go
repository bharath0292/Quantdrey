package brokersentity

import (
	"time"
)

type Broker struct {
	ID          uint64     `gorm:"primaryKey"`
	Name        string     `gorm:"uniqueIndex;not null"`
	DisplayName string     `gorm:"column:display_name"`
	CreatedAt   *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Broker) TableName() string {
	return "brokers"
}
