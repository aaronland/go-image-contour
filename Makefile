GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/contour cmd/contour/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/contour-svg cmd/contour-svg/main.go
