package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/takatoh/randomstring"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
		os.Exit(0)
	}

	rs := randomstring.New("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
	fmt.Println(rs.Rand(n))
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <N>\n", os.Args[0])
}
