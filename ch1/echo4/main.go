package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	iterations := 10 * 1000 * 1000
	sep := " "
	args := ""

	// test joining
	start := time.Now()
	for i := 0; i < iterations; i++ {
		args = strings.Join(os.Args[1:], sep)
	}
	end := time.Since(start).Seconds()
	fmt.Printf("Joining took %.2fs seconds. ", end)

	// test concatenating
	start = time.Now()
	for i := 0; i < iterations; i++ {
		args = ""
		for _, arg := range os.Args[1:] {
			args += sep + arg
		}
	}
	end = time.Since(start).Seconds()
	fmt.Printf("Concatenating took %.2fs seconds. ", end)
}
