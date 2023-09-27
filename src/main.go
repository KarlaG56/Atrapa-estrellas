package main

import (
	"Juego/src/scenes"
	"fmt"
	"fyne.io/fyne/v2/app"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"os"
	"time"
)

func playMusic() {
	file, err := os.Open("src/assets/Beautiful.mp3")
	if err != nil {
		fmt.Println("Error abriendo el archivo:", err)
		return
	}

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		fmt.Println("Error decodificando el archivo:", err)
		return
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
}

func main() {

	go playMusic()

	myApp := app.New()
	scenes.NewMainScene(myApp)
	myApp.Run()
}
