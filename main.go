package main

import (
	"net/http"
	"os"
	"time"

	systray "github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func main() {
	http.DefaultClient.Timeout = 2 * time.Second

	systray.Run(onReady, onExit)
}

func sleep(seconds time.Duration) {
	time.Sleep(seconds * time.Second)
}

func onReady() {
	systray.SetIcon(readIcon())

	var status string
	for {
		status = "❌"
		sleep(2)

		systray.SetTitle(" ...")
		_, err := http.DefaultClient.Get("http://localhost:3001")
		if err == nil {
			status = "✅"
		}
		systray.SetTitle(status)
	}
}

func onExit() {
	// clean up here
}

func readIcon() []byte {
	path := "/home/celso/chatbot/packages/admin/public/static/images/ocelot_avatar.png"

	bytes, err := os.ReadFile(path)
	if err != nil {
		bytes = icon.Data
	}

	return bytes
}
