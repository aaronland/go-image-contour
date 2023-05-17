package contour

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var source_uri string
var target_uri string

var n int
var scale float64

var extra_transformations multi.MultiCSVString

// DefaultFlagSet returns a `flag.FlagSet` instance configured with the default flags
// for running an image contour-ing application.
func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("contour")

	fs.IntVar(&n, "n", 12, "")
	fs.Float64Var(&scale, "scale", 1.0, "")

	fs.StringVar(&source_uri, "source-uri", "file:///", "A valid gocloud.dev/blob.Bucket URI where images are read from.")
	fs.StringVar(&target_uri, "target-uri", "file:///", "A valid gocloud.dev/blob.Bucket URI where images are written to.")
	fs.Var(&extra_transformations, "transformation-uri", "Zero or more additional `transform.Transformation` URIs used to further modify an image after resizing (and before any additional colour profile transformations are performed).")

	return fs
}
