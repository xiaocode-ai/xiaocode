package setup

import (
	"gorm.io/gorm"
	"os"
)

type Setup struct {
	SystemConfigDir string
	SystemDB        *gorm.DB
	ProjectDB       *gorm.DB
}

func New() *Setup {
	return &Setup{
		SystemConfigDir: os.Getenv("HOME") + "/.xiaocode",
	}
}
