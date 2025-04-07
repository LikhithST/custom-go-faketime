package ftime

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var fakeTime time.Time

// Now returns the current time, but if the environment variable
// LD_PRELOAD contains libfaketime then it returns the time
// specified by that library (using date command).
func Now() (t time.Time) {
	var err error
	var env string
	var cmd *exec.Cmd
	var stdout []byte
	var timestamp int64

	if !fakeTime.IsZero() {
		return fakeTime
	}

	if strings.Index(os.Getenv("LD_PRELOAD"), "libfaketime") == -1 {
		currentTime := time.Now()
		fictitiousDelayStr := os.Getenv("FICTITIOUS_DELAY")
		if fictitiousDelayStr != "" {
			delay, err := strconv.ParseFloat(fictitiousDelayStr, 64)
			delay_float := float64(delay)
			currentTime = currentTime.Add(time.Duration(delay_float * float64(time.Second)))
			if err != nil {
				fmt.Println("Error converting FICTITIOUS_DELAY to integer:", err)
				return
			}

		}
		longDistanceConnectionDelay := os.Getenv("LONG_DISTANCE_CONNECTION_DELAY")
		if longDistanceConnectionDelay != "" {
			delay, err := strconv.ParseFloat(longDistanceConnectionDelay, 64)
			delay_float := float64(delay)
			currentTime = currentTime.Add(time.Duration(delay_float * float64(time.Second)))
			if err != nil {
				fmt.Println("Error converting LONG_DISTANCE_CONNECTION_DELAY to integer:", err)
				return
			}

		}
		processDelay := os.Getenv("PROCESS_DELAY")
		if processDelay != "" {
			delay, err := strconv.ParseFloat(processDelay, 64)
			delay_float := float64(delay)
			currentTime = currentTime.Add(time.Duration(delay_float * float64(time.Second)))
			if err != nil {
				fmt.Println("Error converting PROCESS_DELAY to integer:", err)
				return
			}

		}
		return currentTime
	}

	cmd = exec.Command("date", "+%s")
	stdout, err = cmd.Output()
	if err != nil {
		log.Panic(err)
	}

	env = strings.TrimSpace(string(stdout))

	timestamp, err = strconv.ParseInt(env, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	t = time.Unix(timestamp, 0)
	return
}

// SetTime changes the date.
func SetTime(t time.Time) {
	fakeTime = t
}

// ClrTime .
func ClrTime() {
	fakeTime = time.Time{}
}
