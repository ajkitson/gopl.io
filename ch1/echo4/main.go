package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// args := strings.Join(os.Args[0:], " ")
	args := strings.Join(os.Args, " ")
	fmt.Println(args)
}
