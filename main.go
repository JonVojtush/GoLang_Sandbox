package main

import (
	"strconv"
	"syscall/js"
)

var (
	document = js.Global().Get("document")
	// console  = js.Global().Get("console")
	err error
)

func updateUI(this js.Value, args []js.Value) interface{} {
	var (
		aStr string
		a    int
		bStr string
		b    int
	)

	if aStr = document.Call("getElementById", "a").Get("value").String(); aStr == "" {
		document.Call("getElementById", "error").Set("innerText", "box A (left) is empty")
	}
	if bStr = document.Call("getElementById", "b").Get("value").String(); bStr == "" {
		document.Call("getElementById", "error").Set("innerText", "box B (right) is empty")
	}
	if a, err = strconv.Atoi(aStr); err != nil {
		document.Call("getElementById", "error").Set("innerText", "was not able to convert box a to a number")
	}
	if b, err = strconv.Atoi(bStr); err != nil {
		document.Call("getElementById", "error").Set("innerText", "was not able to convert box b to a number")
	}
	result := a + b
	document.Call("getElementById", "result").Set("value", result)
	return nil // no error occurred
}

func main() {
	js.Global().Set("varFromGoToJS", "Hello, I am a variable set from Go, but called from JS.")
	js.Global().Set("updateUI", js.FuncOf(updateUI))
	select {} // a `select` statement at the end of the `main()` function. This is necessary to prevent the Go program from exiting, as the WebAssembly binary will be terminated when the Go program exits.
}
