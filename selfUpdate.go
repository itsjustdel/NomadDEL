package main

import (
	"crypto/ed25519"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"github.com/fynelabs/fyneselfupdate"
	"github.com/fynelabs/selfupdate"
)

// selfManage turns on automatic updates
func selfManage(a fyne.App, w fyne.Window) {
	publicKey := ed25519.PublicKey{190, 196, 251, 69, 245, 240, 238, 252, 245, 111, 179, 246, 172, 145, 46, 234, 116, 36, 31, 199, 207, 86, 68, 118, 226, 153, 128, 113, 147, 8, 77, 51}

	// The public key above matches the signature of the below file served by our CDN
	httpSource := selfupdate.NewHTTPSource(nil, "https://geoffrey-test-artefacts.fynelabs.com/self-update/48/489280fe-faf3-46cf-912f-6344c2dc0030/{{.OS}}-{{.Arch}}/{{.Executable}}{{.Ext}}")

	config := fyneselfupdate.NewConfigWithTimeout(a, w, time.Minute, httpSource, selfupdate.Schedule{FetchOnStart: true, Interval: time.Hour * 12}, publicKey)

	_, err := selfupdate.Manage(config)
	if err != nil {
		log.Println("Error while setting up update manager: ", err)
		return
	}
}
