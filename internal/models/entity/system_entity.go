package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type SystemEntity struct {
	UUID        string    `gorm:"primaryKey;not null;"`               // 系统配置的唯一标识
	Key         string    `gorm:"not null;uniqueIndex"`               // 系统配置的键
	Value       string    `gorm:"not null"`                           // 系统配置的值
	Description string    `gorm:""`                                   // 系统配置的描述
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"` // 创建时间
	UpdatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"` // 更新时间
}

// BeforeUpdate
// 在更新记录之前调用
//
// 设置 UpdatedAt 时间
func (e *SystemEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	e.UpdatedAt = time.Now()
	return nil
}

// BeforeCreate
// 在创建记录之前调用
//
// 设置 UUID、CreatedAt 和 UpdatedAt 时间
func (e *SystemEntity) BeforeCreate(tx *gorm.DB) (err error) {
	e.UUID = uuid.New().String()
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return nil
}
