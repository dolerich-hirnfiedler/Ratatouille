package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type onlyTheme struct {
	fyne.Theme
}

func newTheme() fyne.Theme {
	return &onlyTheme{Theme: theme.DefaultTheme()}
}

func (onlyTheme *onlyTheme) Color(color fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return onlyTheme.Theme.Color(color, theme.VariantLight)
}

func (onlyTheme *onlyTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 12
	}
	return onlyTheme.Theme.Size(name)
}
