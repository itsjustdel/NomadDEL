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
    publicKey := ed25519.PublicKey{241, 170, 13, 63, 95, 130, 43, 112, 245, 238, 209, 140, 85, 10, 139, 35, 11, 104, 200, 115, 170, 254, 141, 101, 172, 169, 16, 90, 207, 152, 115, 213}

    // The public key above matches the signature of the below file served by our CDN
    httpSource := selfupdate.NewHTTPSource(nil, "https://geoffrey-test-artefacts.fynelabs.com/self-update/6e/6e1de1d8-8255-4cbd-8997-33aa35da02e7/{{.OS}}-{{.Arch}}/{{.Executable}}{{.Ext}}")

    config := fyneselfupdate.NewConfigWithTimeout(a, w, time.Minute, httpSource, selfupdate.Schedule{FetchOnStart: true, Interval: time.Hour * 12}, publicKey)

    _, err := selfupdate.Manage(config)
    if err != nil {
        log.Println("Error while setting up update manager: ", err)
        return
    }
}