package gui

import (
	"fyne.io/fyne/v2"
)

type onlyLayout struct {
	top, left, right, content fyne.CanvasObject
}

const sideWidth = 220

// Layout implements fyne.Layout.
func (onlyLayout *onlyLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	topHeigth := onlyLayout.top.MinSize().Height
	onlyLayout.top.Resize(fyne.NewSize(size.Width, topHeigth))
	onlyLayout.left.Move(fyne.NewPos(0, topHeigth))
	onlyLayout.left.Resize(fyne.NewSize(sideWidth, size.Height-topHeigth))

	// log.Println("Size:", size)

}

// MinSize implements fyne.Layout.
func (onlyLayout *onlyLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(10, 10)
}

func newOnlyLayout(top, left, right, content fyne.CanvasObject) fyne.Layout {
	return &onlyLayout{
		top:     top,
		left:    left,
		right:   right,
		content: content,
	}
}
