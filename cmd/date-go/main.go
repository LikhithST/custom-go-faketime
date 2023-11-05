package main

import (
	"fmt"
	"os"
	"time"
	"github.com/harkaitz/go-faketime"
)
func main() {
	var t time.Time
	if len(os.Args) == 1 || os.Args[1] == "ftime.Time" {
		t = ftime.Now()
	} else {
		t = time.Now()
	}
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}
