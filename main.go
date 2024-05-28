package main

import (
	"strconv"
	"syscall/js"
)

var (
	document = js.Global().Get("document")
	console  = js.Global().Get("console")
	// err      error
)

func main() {
	// go startHttpServer() // run the HTTP server in a goroutine so it doesn't block the execution of other JavaScript code
	/*jsErrCh := make(chan error)
	err = <-jsErrCh // wait for JavaScript error to happen
	if err != nil {
		console.Call("error", err)
	}*/
	js.Global().Set("updateUI", js.FuncOf(updateUI))
	js.Global().Set("varFromGoToJS", "Hello, I am a variable set from Go, but called from JS.")
	select {} // a `select` statement at the end of the `main()` function. This is necessary to prevent the Go program from exiting, as the WebAssembly binary will be terminated when the Go program exits.
}

func updateUI(this js.Value, args []js.Value) interface{} {
	aStr := document.Call("getElementById", "a").Get("value").String()
	bStr := document.Call("getElementById", "b").Get("value").String()
	/*if aStr == "" || bStr == "" {
		return errors.New("input fields are empty") // return error if input fields are empty
	}*/
	a, _ := strconv.Atoi(aStr)
	b, _ := strconv.Atoi(bStr)
	result := a + b
	document.Call("getElementById", "result").Set("value", result)
	console.Call("log", "Result updated.")
	return nil // no error occurred
}
