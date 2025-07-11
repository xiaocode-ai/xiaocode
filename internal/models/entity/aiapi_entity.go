package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// AiApiEntity
// AI 的接口代理实体
//
//   - UUID 唯一标识
//   - Link 接口链接
//   - ApiKey 接口密钥
//   - CreatedAt 创建时间
//   - UpdatedAt 更新时间
type AiApiEntity struct {
	UUID        string    `gorm:"primaryKey;not null;"`               // 是 AI 接口代理的唯一标识
	Link        string    `gorm:"not null;uniqueIndex"`               // 链接，必须唯一
	ApiKey      string    `gorm:"not null"`                           // 密钥，用于身份验证
	Description string    `gorm:"not null"`                           // 描述信息
	Active      bool      `gorm:"not null"`                           // 激活状态
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"` // 创建时间，默认当前时间
	UpdatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"` // 更新时间，默认当前时间
}

// BeforeUpdate
// 在创更新记录之前调用
//
// 设置 CreatedAt 时间
func (e *AiApiEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	e.UpdatedAt = time.Now()
	return nil
}

// BeforeCreate
// 在创建记录之前调用
//
// 设置 UUID、CreatedAt 和 UpdatedAt 时间
func (e *AiApiEntity) BeforeCreate(tx *gorm.DB) (err error) {
	e.UUID = uuid.New().String()
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return nil
}
