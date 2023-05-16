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

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("contour")

	fs.IntVar(&n, "n", 12, "")
	fs.Float64Var(&scale, "scale", 1.0, "")

	fs.StringVar(&source_uri, "source-uri", "file:///", "")
	fs.StringVar(&target_uri, "target-uri", "file:///", "")
	fs.Var(&extra_transformations, "transformation-uri", "")

	return fs
}
