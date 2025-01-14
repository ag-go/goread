package category

import (
	"github.com/TypicalAM/goread/internal/theme"
	"github.com/charmbracelet/lipgloss"
)

// popupStyle is the style of the popup window.
type popupStyle struct {
	general             lipgloss.Style
	heading             lipgloss.Style
	list                lipgloss.Style
	choice              lipgloss.Style
	choiceTitle         lipgloss.Style
	choiceDesc          lipgloss.Style
	selectedChoice      lipgloss.Style
	selectedChoiceTitle lipgloss.Style
	selectedChoiceDesc  lipgloss.Style
}

// newPopupStyle creates a new popup style.
func newPopupStyle(colors *theme.Colors, width, height int) popupStyle {
	general := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Width(width - 2).
		Height(height - 2).
		Border(lipgloss.NormalBorder()).
		BorderForeground(colors.Color1)

	heading := lipgloss.NewStyle().
		Margin(1, 0, 1, 0).
		Width(width - 2).
		Align(lipgloss.Center).
		Italic(true)

	list := lipgloss.NewStyle().
		Margin(0, 4).
		Width(width - 2).
		Height(10)

	choice := lipgloss.NewStyle().
		PaddingLeft(1).
		MarginBottom(1).
		Border(lipgloss.RoundedBorder(), false, false, false, true)

	choiceTitle := lipgloss.NewStyle().
		Foreground(colors.Text)

	choiceDesc := lipgloss.NewStyle().
		Foreground(colors.TextDark)

	selectedChoice := choice.Copy().
		Italic(true).
		BorderForeground(colors.Color3)

	selectedChoiceTitle := lipgloss.NewStyle().
		Foreground(colors.Color3)

	selectedChoiceDesc := lipgloss.NewStyle().
		Foreground(colors.Color2)

	return popupStyle{
		general:             general,
		heading:             heading,
		list:                list,
		choice:              choice,
		choiceTitle:         choiceTitle,
		choiceDesc:          choiceDesc,
		selectedChoice:      selectedChoice,
		selectedChoiceTitle: selectedChoiceTitle,
		selectedChoiceDesc:  selectedChoiceDesc,
	}
}
