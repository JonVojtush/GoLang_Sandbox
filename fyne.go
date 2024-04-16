package main

import (
	"log"
	// "net/http" // https://www.digitalocean.com/community/tutorials/how-to-make-http-requests-in-go

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/yosssi/gcss"
)

func GCSS() {
	if _, err := gcss.CompileFile("C:/Users/jonat/go/src/OpenNomad/styles/style.gcss"); err != nil {
		// http.Error(window, err.Error(), http.StatusInternalServerError)
		log.Print(err)
	} /*else {
		http.ServeFile(window, r, cssPath)
	}*/
}

func main() {
	GCSS()
	a := app.New()
	// a.SetIcon()
	w := a.NewWindow("Hanna's Tiny Adventures")
	w.SetContent(
		container.NewVBox(
			widget.NewLabel("Hello Jon!")))
	w.ShowAndRun()
}
