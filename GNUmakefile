.POSIX:
.SUFFIXES:
.PHONY: all clean install check
all:
PROJECT=go-faketime
VERSION=1.0.0
PREFIX=/usr/local

## -- BLOCK:go --
build/date-go$(EXE):
	mkdir -p build
	go build -o $@ $(GO_CONF) ./cmd/date-go
all: build/date-go$(EXE)
install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp build/date-go$(EXE) $(DESTDIR)$(PREFIX)/bin
clean:
	rm -f build/date-go$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp LICENSE $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
