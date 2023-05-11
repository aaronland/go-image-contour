package contour

import (
	"context"
	"image"

	"github.com/aaronland/go-image/transform"
)

type ContourTransformation struct {
	transform.Transformation
}

func init() {
	ctx := context.Background()
	transform.RegisterTransformation(ctx, "contour", NewContourTransformation)
}

func NewContourTransformation(ctx context.Context, str_url string) (transform.Transformation, error) {

	tr := &ContourTransformation{}
	return tr, nil
}

func (tr *ContourTransformation) Transform(ctx context.Context, im image.Image) (image.Image, error) {
	return ContourImage(ctx, im)
}
