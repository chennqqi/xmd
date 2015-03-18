package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wookoouk/xmd/xmd"
)

const (
	zipTmpOutput = "tmp/out/"
)

var usage = `Usage: xmd [options...]

Options:
  init  setup current git repo to use xmd.
  tomd    convert file from docx to xmd.
  tox  convert file from xmd to docx.
`

func main() {

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("no args...")
		usageAndExit("")
	} else {

		if flag.Arg(0) == "init" {
			xmd.SetupGit()

		} else if flag.Arg(0) == "tomd" {
			xmd.ToMD(flag.Arg(1), flag.Arg(2))

		} else if flag.Arg(0) == "tox" {
			xmd.ToX(flag.Arg(1), flag.Arg(2))

		} else {
			usageAndExit("")

		}
	}
}

func usageAndExit(message string) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
