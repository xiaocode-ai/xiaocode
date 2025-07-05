package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
)

func (custom *CustomTui) SetLoading(loading bool) {
	custom.loading = loading
}

func GetCustom() *CustomTui {
	return &CustomTui{
		loading: true,
	}
}

func GetInit(custom *CustomTui) (spinner.Model, textinput.Model) {
	spin := spinner.New()
	spin.Spinner = spinner.MiniDot
	spin.Style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(consts.ColorDarkGreen))

	input := textinput.New()
	input.Focus()

	return spin, input
}
