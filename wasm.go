package main

import (
	_ "errors"
	"fmt"
	"strconv"
	"syscall/js"
)

var (
	document = js.Global().Get("document")
	console  = js.Global().Get("console")
	// server   http.Handler
)

/*func httpServer() {
	const serverRoot = "http/"
	server = http.FileServer(http.Dir(serverRoot))
	log.Print("Serving " + serverRoot + " on http://localhost:3380")
	http.ListenAndServe(":3380", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		server.ServeHTTP(resp, req)
	}))
}*/

func main() {
	// httpServer()

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

	js.Global().Set("varFromGoToJS", "Hello, I ama variable set from Go, but called from JS.")

	select {} // a `select` statement at the end of the `main()` function. This is necessary to prevent the Go program from exiting, as the WebAssembly binary will be terminated when the Go program exits.
}
