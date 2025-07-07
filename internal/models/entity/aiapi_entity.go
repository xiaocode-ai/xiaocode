package entity

import (
	"time"

	"gorm.io/gorm"
)

// AiApiEntity
//
//   - UUID 唯一标识
//   - Limit 限制
//   - ApiKey 接口密钥
//   - CreatedAt 创建时间
//   - UpdatedAt 更新时间
type AiApiEntity struct {
	gorm.Model
	UUID        string    `gorm:"primaryKey"`
	Limit       int       `gorm:"not null;unique"`
	ApiKey      string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Active      bool      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
