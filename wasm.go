package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

var (
	window   = js.Global().Get("window")
	document = js.Global().Get("document")
	console  = js.Global().Get("console")
)

// TASK: set varFromGoToJS VARIABLE FROM GO?

func main() {
	js.Global().Set("goLog", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return "goLog() called with no args."
		}
		return fmt.Sprintf("goLog() called with 1 arg: Hi %s.\n", args[0].String())
	}))

	js.Global().Set("updateUI", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		aStr := document.Call("getElementById", "a").Get("value").String()
		bStr := document.Call("getElementById", "b").Get("value").String()
		a, _ := strconv.Atoi(aStr)
		b, _ := strconv.Atoi(bStr)
		result := a + b
		document.Call("getElementById", "result").Set("value", result)
		console.Call("log", "Result updated.")
		return nil
	}))

	select {} // a `select` statement at the end of the `main()` function. This is necessary to prevent the Go program from exiting, as the WebAssembly binary will be terminated when the Go program exits.
}
