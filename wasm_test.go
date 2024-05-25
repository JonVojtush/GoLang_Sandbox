package main

import (
	"syscall/js"
	"testing"
)

func TestUpdateUI(t *testing.T) {
	// set input values
	aInput := document.Call("getElementById", "a")
	bInput := document.Call("getElementById", "b")
	result := document.Call("getElementById", "result")
	aInput.Set("value", 50)
	bInput.Set("value", 75)
	// call updateUI function
	js.Global().Call("updateUI")
	// test result value is correct
	if result.Get("value").String() != "125" {
		t.Errorf("Expected %v, got %v", 125, result.Get("value"))
	}
}
