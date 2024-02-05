PROJECT=go-faketime
VERSION=1.0.0
PREFIX=/usr/local
all:
clean:
install:

## -- BLOCK:go --
all: all-go
install: install-go
clean: clean-go
deps: deps-go

build/date-go$(EXE): deps
	go build -o $@ $(GO_CONF) ./cmd/date-go

all-go: build/date-go$(EXE)
deps-go:
	@mkdir -p build
install-go:
	@install -d $(DESTDIR)$(PREFIX)/bin
	cp build/date-go$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f build/date-go$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	@mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp LICENSE  $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
