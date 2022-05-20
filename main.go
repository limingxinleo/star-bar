package main

import (
	"encoding/json"
	"fmt"
	"github.com/limingxinleo/star-bar/config"
	"github.com/limingxinleo/star-bar/repo"
	"github.com/limingxinleo/star-bar/voice"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"io"
	"net/http"
	"runtime"
	"time"
)

func main() {
	runtime.LockOSThread()

	cf := config.Init()
	voice.Init()

	var starCount int64 = 0

	cocoa.TerminateAfterWindowsClose = false
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		fmt.Println("Get Status Bar")
		obj := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		fmt.Println("Retain")
		obj.Retain()
		fmt.Println("Set Title")
		obj.Button().SetTitle("GitHub Star")
		go func() {
			for {
				request, _ := http.NewRequest("GET", "https://api.github.com/repos/"+cf.Repo, nil)
				request.Header.Set("Authorization", "Token "+cf.Token)
				response, err := (&http.Client{}).Do(request)
				if err != nil {
					continue
				}

				body, _ := io.ReadAll(response.Body)
				repo := new(repo.Repo)
				err = json.Unmarshal(body, repo)
				if err == nil {
					if starCount < repo.StargazersCount && starCount != 0 {
						go voice.Play()
					}

					starCount = repo.StargazersCount

					core.Dispatch(func() {
						obj.Button().SetTitle(fmt.Sprintf("HF: %d", starCount))
					})
				}

				<-time.After(time.Minute)
			}
		}()
		fmt.Println("New Menu Item")
		itemQuit := cocoa.NSMenuItem_New()
		itemQuit.SetTitle("退出")
		itemQuit.SetAction(objc.Sel("terminate:"))

		fmt.Println("New Menu")
		menu := cocoa.NSMenu_New()
		menu.AddItem(itemQuit)
		obj.SetMenu(menu)
	})
	fmt.Println("Run")
	app.Run()
}
