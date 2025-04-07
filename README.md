A TIME.NOW() THAT RESPECTS FAKETIME(1)
======================================

On UNIX like operating systems there's a useful utility (faketime(1))
that allows executing a program by faking its clock time.

This is done by replacing a C library function with LD_PRELOAD. Of
course go doesn't use the c library so the program fails.

This library implements a copy of "time.Now()" with the name 
"ftime.Now()". It works exactly same way "time.Now()" does
but uses the command "date" to get the date and time when the
LD_PRELOAD environment variable is set.

## Example program

```sh
$ go build -o build/date-go ./cmd/date-go
$ faketime '2040-01-01' ./build/date-go time.Time
2023-11-05 15:48:40
$ faketime '2040-01-01' ./build/date-go ftime.Time
2040-01-01 00:00:00
```

## Go documentation

    package ftime // import "github.com/harkaitz/go-faketime"
    
    func ClrTime()
    func Now() (t time.Time)
    func SetTime(t time.Time)

## custom implementation usage

```
export FICTITIOUS_DELAY=0.250

export LONG_DISTANCE_CONNECTION_DELAY=0.10

export PROCESS_DELAY=0.001
```

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
