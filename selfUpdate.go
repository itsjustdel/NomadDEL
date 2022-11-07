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
	publicKey := ed25519.PublicKey{218, 62, 76, 111, 101, 139, 186, 194, 145, 238, 21, 204, 174, 180, 242, 35, 58, 139, 162, 168, 245, 15, 70, 77, 82, 100, 178, 188, 3, 194, 166, 105}

	// The public key above match the signature of the below file served by our CDN
	httpSource := selfupdate.NewHTTPSource(nil, "https://geoffrey-test-artefacts.fynelabs.com/self-update/3b/3b2ad690-331f-4940-b2ad-6660a0708204/{{.OS}}-{{.Arch}}/{{.Executable}}{{.Ext}}")

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
