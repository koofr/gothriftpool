package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/koofr/gothriftpool"
)

const usage = `gothriftpool [options] <iface>

Go Thrift pool proxy generator.

Examples:

gothriftpool myservice.MyService

Options:
`

func fatal(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func main() {
	var packageName = flag.String("p", "", "generated package name")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
		flag.PrintDefaults()
		os.Exit(2)
	}

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
	}

	iface := flag.Arg(0)

	generator, err := gothriftpool.NewGenerator(iface)
	if err != nil {
		fatal(err)
	}

	if *packageName != "" {
		generator.ProxyPackage = *packageName
	}

	code, err := generator.Generate()

	if err != nil {
		fatal(err)
	}

	os.Stdout.Write(code)
}
