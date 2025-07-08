package xtui

import "github.com/charmbracelet/lipgloss"

func RightConnectNormalTable() lipgloss.Border {
	return lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┬",
		BottomLeft:  "└",
		BottomRight: "┴",
	}
}
