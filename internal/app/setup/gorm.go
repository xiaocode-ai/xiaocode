package setup

import (
	"github.com/xiaocode-ai/xiaocode/internal/models/entity"
	"github.com/xiaocode-ai/xiaocode/pkg/xerr"
	"github.com/xiaocode-ai/xiaocode/pkg/xlog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectSystemDatabase
// 连接系统数据库
//
// 连接系统数据库并进行数据库迁移
func (s *Setup) ConnectSystemDatabase() {
	xlog.Logger(xerr.XLevelDebug, xerr.XTagSetup, xerr.XSUCCESS, "连接系统数据库")

	// 连接数据库
	db, err := gorm.Open(sqlite.Open(s.SystemConfigDir+"/database.db"), &gorm.Config{})
	if err != nil {
		panic("链接数据库失败" + err.Error())
	}

	// 数据库迁移内容保存
	migrateErr := db.AutoMigrate(
		&entity.AiApiEntity{},
		&entity.SystemEntity{},
	)
	if migrateErr != nil {
		xlog.Logger(xerr.XLevelError, xerr.XTagSetup, xerr.XFAIL, "数据库迁移失败: "+migrateErr.Error())
		panic("数据库迁移失败: " + migrateErr.Error())
	}

	s.SystemDB = db
}

// ConnectProjectDatabase
// 连接项目数据库
//
// 连接项目数据库并进行数据库迁移
func (s *Setup) ConnectProjectDatabase() {
	xlog.Logger(xerr.XLevelDebug, xerr.XTagSetup, xerr.XSUCCESS, "连接项目数据库")

	// 连接数据库
	db, err := gorm.Open(sqlite.Open(".xiaocode/database.db"), &gorm.Config{})
	if err != nil {
		panic("链接数据库失败" + err.Error())
	}

	// 数据库迁移内容保存
	migrateErr := db.AutoMigrate()
	if migrateErr != nil {
		xlog.Logger(xerr.XLevelError, xerr.XTagSetup, xerr.XFAIL, "数据库迁移失败: "+migrateErr.Error())
		panic("数据库迁移失败: " + migrateErr.Error())
	}

	s.ProjectDB = db
}
