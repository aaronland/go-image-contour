//go:build wasmjs
package main

import (
	"log"
	"syscall/js"

	"github.com/aaronland/go-image-contour/v2/wasm"
)

func main() {

	contour_func := wasm.ContourFunc()
	defer contour_func.Release()

	js.Global().Set("contour", contour_func)

	c := make(chan struct{}, 0)

	log.Println("contour function initialized")
	<-c
}
