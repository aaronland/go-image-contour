//go:build wasmjs
package wasm

import (
	"bytes"
	"bufio"
	"context"
	"fmt"
	"syscall/js"
	"encoding/base64"
	"image"
	_ "image/jpeg"
	_ "image/gif"	
	_ "image/png"
	
	"github.com/aaronland/go-image-contour/v2"
)

func ContourSVGFunc() js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		im_b64 := args[0].String()
		n := args[1].Int()
		scale := 1.0
		
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			ctx := context.Background()
			
			resolve := args[0]
			reject := args[1]

			im_b, err := base64.StdEncoding.DecodeString(im_b64)

			if err != nil {
				reject.Invoke(fmt.Printf("Failed to decode b64 image, %v\n", err))
				return nil
			}

			im_r := bytes.NewReader(im_b)
			im, _, err := image.Decode(im_r)

			if err != nil {
				reject.Invoke(fmt.Printf("Failed to decode image, %v\n", err))
				return nil
			}

			var buf bytes.Buffer
			wr := bufio.NewWriter(&buf)
			
			err = contour.ContourImageSVG(ctx, wr, im, n, scale)

			if err != nil {
				reject.Invoke(fmt.Printf("Failed to contour feature, %v\n", err))
				return nil
			}

			wr.Flush()

			resolve.Invoke(buf.String())
			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
	
}
