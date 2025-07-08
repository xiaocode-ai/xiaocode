package setup

import (
	"github.com/pelletier/go-toml/v2"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
	"github.com/xiaocode-ai/xiaocode/internal/models/dto"
	"github.com/xiaocode-ai/xiaocode/pkg/xerr"
	"github.com/xiaocode-ai/xiaocode/pkg/xlog"
	"os"
)

// CheckAndCreateSystemProfile
// 检查并创建系统配置文件夹
//
// 如果系统配置文件夹不存在，则创建系统配置文件夹
func (s *Setup) CheckAndCreateSystemProfile() {
	xlog.Logger(xerr.XLevelDebug, xerr.XTagSetup, xerr.XSUCCESS, "检查系统配置文件夹是否存在")
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
	xlog.Logger(xerr.XLevelDebug, xerr.XTagSetup, xerr.XSUCCESS, "检查项目配置文件夹是否存在")
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

// SystemProfileLoad
// 加载系统配置文件
//
// 加载系统配置文件，将配置内容从 toml 文件解析并加载到全局配置环境中。
func (s *Setup) SystemProfileLoad() {
	xlog.Logger(xerr.XLevelDebug, xerr.XTagSetup, xerr.XSUCCESS, "加载系统配置文件")
	// 获取 toml 文件并加载进入环境中
	readFile, osErr := os.ReadFile(s.SystemConfigDir + "/config.toml")
	if osErr != nil {
		panic(osErr)
	}
	// 解析 toml 文件
	var systemConfig *dto.SystemConfigDTO
	tomlErr := toml.Unmarshal(readFile, &systemConfig)
	if tomlErr != nil {
		panic(tomlErr)
	}
	// 数据读入环境中
	consts.GlobalSystemProfile = systemConfig
}

func (s *Setup) ProjectProfileLoad() {
	xlog.Logger(xerr.XLevelDebug, xerr.XTagSetup, xerr.XSUCCESS, "加载项目配置文件")
	// 获取 toml 文件并加载进入环境中
	readFile, osErr := os.ReadFile(".xiaocode/config.toml")
	if osErr != nil {
		panic(osErr)
	}
	// 解析 toml 文件
	var projectConfig *dto.ProjectConfigDTO
	tomlErr := toml.Unmarshal(readFile, &projectConfig)
	if tomlErr != nil {
		panic(tomlErr)
	}
	// 数据读入环境中
	consts.GlobalProjectProfile = projectConfig
}
