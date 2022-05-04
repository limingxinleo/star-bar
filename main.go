package main

import (
	"encoding/json"
	"fmt"
	"github.com/limingxinleo/star-bar/config"
	"github.com/limingxinleo/star-bar/repo"
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

	cocoa.TerminateAfterWindowsClose = false
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		obj := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		obj.Retain()
		obj.Button().SetTitle("GitHub Star")

		go func() {
			for {
				request, _ := http.NewRequest("GET", "https://api.github.com/repos/"+config.Repo, nil)
				request.Header.Set("Authorization", "Token "+config.Token)
				response, err := (&http.Client{}).Do(request)
				if err != nil {
					continue
				}

				body, _ := io.ReadAll(response.Body)
				repo := new(repo.Repo)
				err = json.Unmarshal(body, repo)
				fmt.Println(repo)
				if err == nil {
					core.Dispatch(func() {
						obj.Button().SetTitle(fmt.Sprintf("HF: %d", repo.StargazersCount))
					})
				}

				time.Sleep(time.Minute)
			}
		}()
		itemQuit := cocoa.NSMenuItem_New()
		itemQuit.SetTitle("退出")
		itemQuit.SetAction(objc.Sel("terminate:"))

		menu := cocoa.NSMenu_New()
		menu.AddItem(itemQuit)
		obj.SetMenu(menu)

	})
	app.Run()
}
