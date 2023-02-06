package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/takatoh/randomstring"
)

const (
	progName    = "strrand"
	progVersion = "v0.4.0"

	uppercases = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercases = "abcdefghijklmnopqrstuvwxyz"
	digits     = "1234567890"
)

func main() {
	flag.Usage = usage
	opt_upper := flag.Bool("upper", false, "Include upper case letters.")
	opt_lower := flag.Bool("lower", false, "Include lower case letters.")
	opt_digit := flag.Bool("digit", false, "Include digits.")
	opt_include := flag.String("include", "", "Include specified letters.")
	opt_squote := flag.Bool("squote", false, "Include single quote(').")
	opt_dquote := flag.Bool("dquote", false, "Include double quote(\").")
	opt_uuid := flag.Bool("uuid", false, "Generate random string based on UUID version 4.")
	opt_version := flag.Bool("version", false, "Show version.")
	flag.Parse()

	if *opt_version {
		fmt.Println(progVersion)
		os.Exit(0)
	}

	if *opt_uuid {
		sbu := randomstring.NewStringBasedUUID()
		s := sbu.Generate()
		fmt.Println(s)
		os.Exit(0)
	}

	pool := ""
	if *opt_upper {
		pool = pool + uppercases
	}
	if *opt_lower {
		pool = pool + lowercases
	}
	if *opt_digit {
		pool = pool + digits
	}
	if len(pool) == 0 {
		pool = uppercases + lowercases + digits
	}
	pool = pool + *opt_include
	if *opt_squote {
		pool = pool + "'"
	}
	if *opt_dquote {
		pool = pool + "\""
	}

	if len(flag.Args()) < 1 {
		usage()
		os.Exit(0)
	}

	n, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		usage()
		os.Exit(0)
	}

	rs := randomstring.New(pool)
	fmt.Println(rs.Rand(n))
}

func usage() {
	fmt.Fprintf(os.Stderr,
		`%s - Generate a random string

Usage:

Random string of length <N>
  %s [options] <N>

Or string based on UUID version 4
  %s -uuid

Options:
`, progName, progName, progName)
	flag.PrintDefaults()
}
