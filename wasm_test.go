// https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/
package main

import (
	_ "testing"

	_ "github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/require"
)

// func testSomething() {}
