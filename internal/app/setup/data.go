package setup

import (
	"errors"
	"github.com/xiaocode-ai/xiaocode/internal/models/entity"
	"github.com/xiaocode-ai/xiaocode/pkg/xerr"
	"github.com/xiaocode-ai/xiaocode/pkg/xlog"
	"gorm.io/gorm"
)

// SystemDatabaseDataPrepare
// 准备系统表数据
//
// 检查系统表数据是否存在，如果不存在则插入默认数据
func (s *Setup) SystemDatabaseDataPrepare() {
	xlog.Logger(xerr.XLevelDebug, xerr.XTagSetup, xerr.XSUCCESS, "准备系统数据库数据")

	db := s.SystemDB

	checkAndInsertData(db, "enable_mcp", "true", "是否开启系统的 MCP 服务")
}

// checkAndInsertData
// 检查数据是否存在，如果不存在则插入新数据
//
//   - key: 数据的键
//   - val: 数据的值
//   - desc: 数据的描述
func checkAndInsertData(db *gorm.DB, key, val, desc string) string {
	var getEntity *entity.SystemEntity
	txGet := db.Where("key = ?", key).First(&getEntity)
	if errors.Is(txGet.Error, gorm.ErrRecordNotFound) {
		// 如果没有找到记录，则插入新记录
		newEntity := entity.SystemEntity{
			Key:         key,
			Value:       val,
			Description: desc,
		}
		txCreate := db.Create(&newEntity)
		if txCreate.Error != nil {
			panic("数据插入失败: " + txCreate.Error.Error())
		}
		return newEntity.Value
	} else if txGet.Error != nil {
		panic("查询数据失败: " + txGet.Error.Error())
	} else {
		return getEntity.Value
	}
}
