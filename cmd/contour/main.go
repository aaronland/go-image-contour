package main

import (
	"context"
	"log"

	"github.com/aaronland/go-image-contour/app/contour"
	_ "github.com/aaronland/go-image/common"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := contour.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to contour images, %v", err)
	}
}
