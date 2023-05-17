# go-image-contour

Opinionated Go package for working with the `fogleman/contourmap` package.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/aaronland/go-image-contour.svg)](https://pkg.go.dev/github.com/aaronland/go-image-contour)

## Tools

$> make cli
go build -mod vendor -ldflags="-s -w" -o bin/contour cmd/contour/main.go
go build -mod vendor -ldflags="-s -w" -o bin/contour-svg cmd/contour-svg/main.go

### contour

![](fixtures/tokyo.jpg)

![](fixtures/tokyo-contour-3.jpg)

### contour-svg

## See also

* https://github.com/aaronland/go-image
* https://github.com/fogleman/contourmap
* https://github.com/fogleman/gg