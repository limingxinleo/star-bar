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

	var starCount uint64 = 0
	title := "GitHub"
	toolTip := ""

	cocoa.TerminateAfterWindowsClose = false
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		fmt.Println("Get Status Bar")
		obj := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		fmt.Println("Retain")
		obj.Retain()
		fmt.Println("Set Title")
		obj.Button().SetTitle(title)
		logo := cocoa.NSImage_InitWithURL(core.NSURL_Init("https://kycmd-pub.knowyourself.cc/github/logo-standard.png"))
		logo.SetSize(core.NSMakeSize(18, 18))
		obj.Button().SetImage(logo)
		obj.Button().SetImagePosition(2)
		nextListen := make(chan bool)
		key := 0
		go func() {
			for {
				request, _ := http.NewRequest("GET", "https://api.github.com/repos/"+cf.Repo, nil)
				request.Header.Set("Authorization", "Token "+cf.Token)
				response, err := (&http.Client{Timeout: time.Second * 5}).Do(request)
				if err != nil {
					fmt.Println(err)
					continue
				}

				body, _ := io.ReadAll(response.Body)
				repo := &repo.Repo{}
				err = json.Unmarshal(body, repo)
				if err == nil {
					if starCount < repo.StargazersCount && starCount != 0 {
						go voice.Play()
					}

					starCount = repo.StargazersCount

					switch key % 2 {
					case 0:
						title = fmt.Sprintf("%d", starCount)
						toolTip = "关注数"
						break
					case 1:
						title = fmt.Sprintf("%d", repo.OpenIssuesCount)
						toolTip = "问题数"
					}

					core.Dispatch(func() {
						obj.Button().SetTitle(title)
						obj.Button().SetToolTip(toolTip)
					})
				}

				select {
				case <-nextListen:
					key++
				case <-time.After(time.Minute):
				}
			}
		}()
		fmt.Println("New Menu Item")
		itemQuit := cocoa.NSMenuItem_New()
		itemQuit.SetTitle("退出")
		itemQuit.SetAction(objc.Sel("terminate:"))

		issueItem := cocoa.NSMenuItem_New()
		issueItem.SetTitle("更换")
		issueItem.SetAction(objc.Sel("nextListen:"))
		cocoa.DefaultDelegateClass.AddMethod("nextListen:", func(_ objc.Object) {
			nextListen <- true
		})

		fmt.Println("New Menu")
		menu := cocoa.NSMenu_New()
		menu.AddItem(issueItem)
		menu.AddItem(itemQuit)
		obj.SetMenu(menu)
	})
	fmt.Println("Run")
	app.Run()
}
