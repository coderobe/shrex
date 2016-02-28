package main

import (
	"fmt"

	"github.com/getlantern/systray"
	"github.com/shex-project/shex/icon"
)

func main() {
	systray.Run(onReady)
	fmt.Println("this is shex")
}

func onReady() {
	// systray init
	systray.SetIcon(icon.Data) // do we need an icon for this to work?
	systray.SetTitle("shex")
	systray.SetTooltip("this is the shex systray. woo")
	// menu items
	mScreenshotDesktop := systray.AddMenuItem("Screenshot entire desktop", "")
	mSettings := systray.AddMenuItem("Settings", "Change some stuff")
	mQuit := systray.AddMenuItem("Quit", "Exit this crap")

	// do we need different goroutines for each menu item?

	go func() {
		<-mScreenshotDesktop.ClickedCh
		// TODO: automagically close systray and take a screenshot
		fmt.Println("ScreenshotDesktop") // TODO: remove debug output
	}()

	go func() {
		<-mSettings.ClickedCh
		// TODO: open settings dialog
		//   - can change keybinds
		//   - can change upload server
		//   - other settings?
		fmt.Println("SettingsDialog") // TODO: remove debug output
	}()

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
		fmt.Println("EXITS WITH SOME SICK KICKFLIPS BRAH")
	}()
}
