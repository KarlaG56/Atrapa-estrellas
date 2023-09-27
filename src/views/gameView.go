package views

import (
	"Juego/src/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"math/rand"
	"time"
)

type GameView struct {
	game       *models.Game
	window     fyne.Window
	scoreLabel *widget.Label
	timerLabel *widget.Label
	stars      []fyne.CanvasObject
	objects    []fyne.CanvasObject
	timeLeft   int
	bgImage    *canvas.Image
}

func NewGameView(app fyne.App) *GameView {
	game := models.NewGame()
	view := &GameView{
		game:       game,
		window:     app.NewWindow("Atrapa las Estrellas"),
		scoreLabel: widget.NewLabel("Score: 0"),
		timerLabel: widget.NewLabel("Time: 60"),
		timeLeft:   60,
	}
	bgImageData, err := ioutil.ReadFile("src/assets/Fondo.jpg")
	if err != nil {
		fmt.Println("Error al cargar la imagen de fondo:", err)
	} else {
		bgResource := fyne.NewStaticResource("background.png", bgImageData)
		view.bgImage = canvas.NewImageFromResource(bgResource)
		view.bgImage.FillMode = canvas.ImageFillStretch
		view.bgImage.Resize(fyne.NewSize(1280, 720))
		view.objects = append([]fyne.CanvasObject{view.bgImage}, view.objects...)
	}
	view.window.Resize(fyne.NewSize(800, 600))

	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), view.scoreLabel, view.timerLabel)

	view.objects = append(view.objects, container)

	return view
}

func (v *GameView) generateStar() {
	imgData, err := ioutil.ReadFile("src/assets/star.png")
	if err != nil {
		fmt.Println("Error al cargar la imagen:", err)
		return
	}
	imgResource := fyne.NewStaticResource("star.png", imgData)
	var starButton *widget.Button
	starButton = widget.NewButtonWithIcon("", imgResource, func() {
		v.game.Score++
		v.scoreLabel.SetText(fmt.Sprintf("Score: %d", v.game.Score))
		v.removeStar(starButton)
	})
	starButton.Resize(fyne.NewSize(40, 40))
	xPos := rand.Float32() * float32(v.window.Canvas().Size().Width-30)
	starButton.Move(fyne.NewPos(float32(xPos), 0))

	v.stars = append(v.stars, starButton)
	v.objects = append(v.objects, starButton)
}

func (v *GameView) removeStar(obj fyne.CanvasObject) {
	// Remove the star or button from stars slice
	for i, s := range v.stars {
		if s == obj {
			v.stars = append(v.stars[:i], v.stars[i+1:]...)
			break
		}
	}

	// Remove the star or button from objects slice
	for i, o := range v.objects {
		if o == obj {
			v.objects = append(v.objects[:i], v.objects[i+1:]...)
			break
		}
	}
}

func (v *GameView) moveStars() {
	for _, star := range v.stars {
		position := star.Position()
		star.Move(fyne.NewPos(position.X, position.Y+5))

		if position.Y > float32(v.window.Canvas().Size().Height) {
			v.game.Score++
			v.scoreLabel.SetText(fmt.Sprintf("Score: %d", v.game.Score))
			v.removeStar(star)
		}
	}
}

func (v *GameView) countdown() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			v.timeLeft--
			v.timerLabel.SetText(fmt.Sprintf("Time: %d", v.timeLeft))
			if v.timeLeft == 0 {
				ticker.Stop()
				v.endGame()
			}
		}
	}()
}

func (v *GameView) endGame() {
	v.window.SetContent(widget.NewLabel("Game Over! Final Score: " + fmt.Sprintf("%d", v.game.Score)))
}

func (v *GameView) Show() {
	content := container.NewWithoutLayout(v.objects...)
	v.window.SetContent(content)
	v.countdown()
	starTicker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for range starTicker.C {
			if v.timeLeft == 0 {
				starTicker.Stop()
				return
			}
			v.generateStar()
			v.moveStars()
			content := container.NewWithoutLayout(v.objects...)
			v.window.SetContent(content)
		}
	}()
	v.window.ShowAndRun()
}
