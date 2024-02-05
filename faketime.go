package ftime

import (
	"time"
	"os/exec"
	"strings"
	"strconv"
	"os"
	"log"
)

var fakeTime time.Time

// Now returns the current time, but if the environment variable
// LD_PRELOAD contains libfaketime then it returns the time
// specified by that library (using date command).
func Now() (t time.Time) {
	var err       error
	var env       string
	var cmd      *exec.Cmd
	var stdout  []byte
	var timestamp int64
	
	if !fakeTime.IsZero() {
		return fakeTime
	}
	
	if strings.Index(os.Getenv("LD_PRELOAD"), "libfaketime") == -1 {
		return time.Now()
	}
	
	cmd = exec.Command("date", "+%s")
	stdout, err = cmd.Output()
	if err != nil { log.Panic(err) }
	
	env = strings.TrimSpace(string(stdout))
	
	timestamp, err = strconv.ParseInt(env, 10, 64)
	if err != nil { log.Panic(err) }
	
	t = time.Unix(timestamp, 0)
	return
}

// SetTime changes the date.
func SetTime(t time.Time) {
	fakeTime = t
}

// ClrTime .
func ClrTime() {
	fakeTime = time.Time {}
}
