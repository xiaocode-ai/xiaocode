package system

import (
	"github.com/xiaocode-ai/xiaocode/internal/app/setup"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SystemDAO struct {
	DB *gorm.DB
}

func (s *SystemDAO) PrepareSystemGorm(setup *setup.Setup) {
	db, err := gorm.Open(sqlite.Open(setup.SystemConfigDir+"/database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	s.DB = db
}