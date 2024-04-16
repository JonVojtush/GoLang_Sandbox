package main

import (
	"fmt"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/yosssi/gcss"
)

var (
	w http.ResponseWriter
	r *http.Request
)

// Pre-process the .gcss stylesheet into .css
func GCSS() {
	cssPath, err := gcss.CompileFile("web/styles/style.gcss")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.ServeFile(w, r, cssPath)
}

func main() {
	GCSS()

	err := app.GenerateStaticWebsite("", &app.Handler{
		Name:  "Github Pages Hello",
		Title: "Github Pages Hello",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("static website generated")
}
