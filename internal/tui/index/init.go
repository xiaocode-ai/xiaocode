package index

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/xiaocode-ai/xiaocode/internal/consts"
)

func (custom *CustomOperate) SetLoading(loading bool) {
	custom.loading = loading
}

func NewCustom() *CustomOperate {
	return &CustomOperate{
		loading: true,
	}
}

func NewKeyboard() *CustomKeyboard {
	return &CustomKeyboard{
		esc: false,
	}
}

func GetInit(custom *CustomOperate) (spinner.Model, textinput.Model) {
	spin := spinner.New()
	spin.Spinner = spinner.MiniDot
	spin.Style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(consts.ColorDarkGreen))

	input := textinput.New()
	input.Focus()

	return spin, input
}

func NewTui(custom *CustomOperate, keyboard *CustomKeyboard) *Tui {
	spin, input := GetInit(custom)

	return &Tui{
		custom:    custom,
		keyboard:  keyboard,
		spinner:   spin,
		textInput: input,
	}
}