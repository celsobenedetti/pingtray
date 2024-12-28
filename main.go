package main

import (
	"net/http"
	"os"
	"time"

	systray "github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

const (
	requestTimeout = 2 * time.Second
	updateInterval = 2 * time.Second
	successStaus   = "✅"
	failureStatus  = "❌"
	loadingStatus  = "..."
)

var (
	iconPath = os.Getenv("SYSTRAY_ICON") // .ico or .png
	endpoint = os.Getenv("SYSTRAY_ENDPOINT")
	client   = &http.Client{Timeout: requestTimeout}
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(readIconFromFS(iconPath))
	quitButton := systray.AddMenuItem("Quit", "quit")
	fetchAndUpdate()

	ticker := time.NewTicker(updateInterval)

	for {
		select {
		case <-ticker.C:
			fetchAndUpdate()
		case <-quitButton.ClickedCh:
			systray.Quit()
		}
	}
}

func fetchAndUpdate() {
	systray.SetTitle(loadingStatus)
	if _, err := client.Get(endpoint); err != nil {
		systray.SetTitle(failureStatus)
		return
	}
	systray.SetTitle(successStaus)
}

func readIconFromFS(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		bytes = icon.Data // Default icon
	}
	return bytes
}

func onExit() {}
