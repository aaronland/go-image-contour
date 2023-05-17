package svg

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/aaronland/go-image-contour"
	"github.com/aaronland/go-image/decode"
	"github.com/aaronland/gocloud-blob/bucket"
	"github.com/sfomuseum/go-flags/flagset"
	"gocloud.dev/blob"
)

// RunOptions is a struct containing configuration details for running an image transformation application.
type RunOptions struct {
	// SourceURI is a `gocloud.dev/blob.Bucket` URI specifying the location where images are read from.
	SourceURI string
	// SourceURI is a `gocloud.dev/blob.Bucket` URI specifying the location where images are written to.
	TargetURI string
	// Logger is a `log.Logger` instance used for logging messages and feedback.
	Logger *log.Logger
}

// Run invokes the application to generate SVG derived from an image's contours using the default flags.
func Run(ctx context.Context, logger *log.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

// Run invokes the application to generate SVG derived from an image's contours using a custom `flag.FlagSet` instance.
func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *log.Logger) error {

	flagset.Parse(fs)

	opts := &RunOptions{
		SourceURI: source_uri,
		TargetURI: source_uri,
		Logger:    logger,
	}

	paths := fs.Args()

	return RunWithOptions(ctx, opts, paths...)
}

// Run invokes the application to generate SVG derived from an image's contours configured using 'opts'.
func RunWithOptions(ctx context.Context, opts *RunOptions, paths ...string) error {

	source_b, err := bucket.OpenBucket(ctx, opts.SourceURI)

	if err != nil {
		return fmt.Errorf("Failed to open source, %w", err)
	}

	defer source_b.Close()

	target_b, err := bucket.OpenBucket(ctx, opts.TargetURI)

	if err != nil {
		return fmt.Errorf("Failed to open target, %w", err)
	}

	defer target_b.Close()

	done_ch := make(chan bool)
	err_ch := make(chan error)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, key := range paths {

		go func(key string) {

			err := deriveSVG(ctx, opts, source_b, target_b, key)

			if err != nil {
				err_ch <- err
			}

			done_ch <- true
		}(key)
	}

	remaining := len(paths)

	for remaining > 0 {
		select {
		case <-ctx.Done():
			return nil
		case <-done_ch:
			remaining -= 1
		case err := <-err_ch:
			return err
		}
	}

	return nil
}

func deriveSVG(ctx context.Context, opts *RunOptions, source_b *blob.Bucket, target_b *blob.Bucket, key string) error {

	if opts.SourceURI == "file:///" {

		abs_key, err := filepath.Abs(key)

		if err != nil {
			return fmt.Errorf("Failed to derive absolute path for %s, %w", key, err)
		}

		key = abs_key
	}

	r, err := bucket.NewReadSeekCloser(ctx, source_b, key, nil)

	if err != nil {
		return fmt.Errorf("Failed to open %s for reading, %v", key, err)
	}

	defer r.Close()

	dec, err := decode.NewDecoder(ctx, key)

	if err != nil {
		return fmt.Errorf("Failed to create decoder for %s, %w", key, err)
	}

	im, _, err := dec.Decode(ctx, r)

	if err != nil {
		return fmt.Errorf("Failed to decode %s, %v", key, err)
	}

	old_ext := filepath.Ext(key)
	new_ext := ".svg"

	new_key := key
	new_key = strings.Replace(new_key, old_ext, "", 1)

	new_key = fmt.Sprintf("%s-%03d%s", new_key, n, new_ext)

	wr, err := target_b.NewWriter(ctx, new_key, nil)

	if err != nil {
		return fmt.Errorf("Failed to create new writer for %s, %w", new_key, err)
	}

	err = contour.ContourImageSVG(ctx, wr, im, n, scale)

	if err != nil {
		return fmt.Errorf("Failed to derive SVG from %s, %w", key, err)
	}

	err = wr.Close()

	if err != nil {
		return fmt.Errorf("Failed to close writer for %s, %w", new_key, err)
	}

	return nil
}
