//go:build wasmjs
package main

import (
	"log"
	"syscall/js"

	"github.com/aaronland/go-image-contour/v2/wasm"
)

func main() {

	contour_func := wasm.ContourImageFunc()
	defer contour_func.Release()

	contour_svg_func := wasm.ContourSVGFunc()
	defer contour_svg_func.Release()
	
	js.Global().Set("contour", contour_func)
	js.Global().Set("contour_svg", contour_svg_func)	

	c := make(chan struct{}, 0)

	log.Println("contour function initialized")
	<-c
}
