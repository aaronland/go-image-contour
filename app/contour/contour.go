package contour

import (
	"context"
	"flag"
	"fmt"
	"log"

	_ "github.com/aaronland/go-image-contour"
	"github.com/aaronland/go-image/app/transform"
	"github.com/sfomuseum/go-flags/flagset"
)

func Run(ctx context.Context, logger *log.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *log.Logger) error {

	flagset.Parse(fs)

	contour_uri := fmt.Sprintf("contour://?n=%d&scale=%f", n, scale)
	suffix := fmt.Sprintf("-contour-%d", n)

	transformation_uris := []string{
		contour_uri,
	}

	for _, e := range extra_transformations {
		transformation_uris = append(transformation_uris, e)
	}

	opts := &transform.RunOptions{
		TransformationURIs: transformation_uris,
		ApplySuffix:        suffix,
		SourceURI:          source_uri,
		TargetURI:          target_uri,
		Logger:             logger,
	}

	paths := fs.Args()

	return transform.RunWithOptions(ctx, opts, paths...)
}
