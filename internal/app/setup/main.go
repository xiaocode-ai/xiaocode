package setup

import "os"

type Setup struct {
	SystemConfigDir string
}

func New() *Setup {
	return &Setup{
		SystemConfigDir: os.Getenv("HOME") + "/.xiaocode",
	}
}
