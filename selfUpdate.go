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
	publicKey := ed25519.PublicKey{161, 4, 167, 35, 193, 51, 72, 186, 86, 233, 179, 89, 147, 137, 164, 12, 53, 124, 98, 50, 206, 154, 215, 164, 83, 223, 126, 48, 11, 178, 165, 58}

	// The public key above matches the signature of the below file served by our CDN
	httpSource := selfupdate.NewHTTPSource(nil, "https://geoffrey-artefacts.fynelabs.com/self-update/b0/b0ffdd83-729e-48e3-8013-53ddc446fec7/{{.OS}}-{{.Arch}}/{{.Executable}}{{.Ext}}")

	config := fyneselfupdate.NewConfigWithTimeout(a, w, time.Minute, httpSource, selfupdate.Schedule{FetchOnStart: true, Interval: time.Hour * 12}, publicKey)

	_, err := selfupdate.Manage(config)
	if err != nil {
		log.Println("Error while setting up update manager: ", err)
		return
	}
}
