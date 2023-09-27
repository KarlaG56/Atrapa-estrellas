package scenes

import (
	"Juego/src/views"
	"fyne.io/fyne/v2"
)

func NewMainScene(app fyne.App) {
	mainView := views.NewGameView(app)
	mainView.Show()
}
