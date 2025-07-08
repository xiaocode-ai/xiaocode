package setup

import (
	"os"
)

// CheckAndCreateSystemProfile
// 检查并创建系统配置文件夹
//
// 如果系统配置文件夹不存在，则创建系统配置文件夹
func (s *Setup) CheckAndCreateSystemProfile() {
	// 检查系统配置文件夹是否存在
	if _, err := os.Stat(s.SystemConfigDir); os.IsNotExist(err) {
		osErr := os.MkdirAll(s.SystemConfigDir, 0755)
		if osErr != nil {
			panic(osErr)
		}
	}

	// 检查配置文件是否存在
	configFile := s.SystemConfigDir + "/config.toml"
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		_, osErr := os.Create(configFile)
		if osErr != nil {
			panic(osErr)
		}
	}
	// 创建数据库文件
	dbFile := s.SystemConfigDir + "/database.db"
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		_, osErr := os.Create(dbFile)
		if osErr != nil {
			panic(osErr)
		}
	}
}

// CheckAndCreateProjectProfile
// 检查并创建项目配置文件夹
//
// 如果项目配置文件夹不存在，则创建项目配置文件夹
func (s *Setup) CheckAndCreateProjectProfile() {
	// 检查当前目录下是否存在 .xiaocode 文件夹
	if _, err := os.Stat(".xiaocode"); os.IsNotExist(err) {
		osErr := os.MkdirAll(".xiaocode", 0755)
		if osErr != nil {
			panic(osErr)
		}
	}

	// 检查配置文件是否存在
	configFile := ".xiaocode/config.toml"
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		_, osErr := os.Create(configFile)
		if osErr != nil {
			panic(osErr)
		}
	}
	// 创建数据库文件
	dbFile := ".xiaocode/database.db"
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		_, osErr := os.Create(dbFile)
		if osErr != nil {
			panic(osErr)
		}
	}
}
