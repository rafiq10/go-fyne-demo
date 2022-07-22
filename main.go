package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Hello world")
	window.Resize(fyne.NewSize(600, 500))
	label := widget.NewLabel("New Content")
	vBox := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		label,
		layout.NewSpacer())
	window.SetContent(vBox)
	label.Wrapping = fyne.TextWrapBreak
	client = &http.Client{Timeout: time.Second * 10}

	go func() {
		for {
			time.Sleep(time.Second * 5)
			fact, err := randFact()
			if err == nil {
				label.SetText(fact)
			}
		}
	}()
	window.ShowAndRun()
}

var client *http.Client

func randFact() (string, error) {
	resp, err := client.Get("https://uselessfacts.jsph.pl/random.json?language=en")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	randomFact, err := decodeRandomFact(resp.Body)
	if err != nil {
		log.Println("error decoding response body", err)
		return "", err
	}
	return randomFact, nil
}

func decodeRandomFact(jsonString io.ReadCloser) (string, error) {
	body := struct {
		Text string `json:"text"`
	}{}

	err := json.NewDecoder(jsonString).Decode(&body)
	if err != nil {
		log.Println("error decoding response body", err)
		return "", err
	}
	return body.Text, nil
}
