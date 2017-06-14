package main

import (
	"fmt"
	"os"
)

func main() {
	for ix, arg := range os.Args[1:] {
		fmt.Println(ix, arg)
	}
}
