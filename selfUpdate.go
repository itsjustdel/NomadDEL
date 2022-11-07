package main

import (
	"crypto/ed25519"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"github.com/fynelabs/fyneselfupdate"
	"github.com/fynelabs/selfupdate"
)

// selfManage turns on automatic update
func selfManage(a fyne.App, w fyne.Window) {
	publicKey := ed25519.PublicKey{161, 155, 249, 197, 223, 115, 158, 168, 245, 147, 70, 61, 243, 28, 161, 67, 217, 25, 40, 5, 16, 114, 113, 88, 85, 232, 116, 119, 200, 198, 57, 212}

	// The public key above match the signature of the below file served by our CDN
	httpSource := selfupdate.NewHTTPSource(nil, "https://geoffrey-test-artefacts.fynelabs.com/self-update/7e/7e42c9e2-3a3d-46b3-a874-e543c1dd45b6/{{.OS}}-{{.Arch}}/{{.Executable}}{{.Ext}}")

	config := fyneselfupdate.NewConfigWithTimeout(a, w, time.Duration(1)*time.Minute,
		httpSource,
		selfupdate.Schedule{FetchOnStart: true, Interval: time.Hour * time.Duration(12)},
		publicKey)

	_, err := selfupdate.Manage(config)
	if err != nil {
		log.Println("Error while setting up update manager: ", err)
		return
	}
}
