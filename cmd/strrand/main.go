package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/takatoh/randomstring"
)

const (
	progName = "strrand"
	progVersion = "v0.1.0"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
`Usage:
  %s [options] <N>

Options:
`, progName)
		flag.PrintDefaults()
	}
	opt_version := flag.Bool("version", false, "Show version.")
	flag.Parse()
	if *opt_version {
		fmt.Println(progVersion)
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
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
