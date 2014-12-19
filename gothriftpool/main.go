package main

import (
	"flag"
	"fmt"
	"github.com/koofr/gothriftpool"
	"io/ioutil"
	"os"
	"path/filepath"
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
	var write = flag.Bool("w", false, "write generated template to $GOPATH/packageName")

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

	if *write {
		parent := filepath.Join(os.Getenv("GOPATH"), "src", generator.ProxyPackage)

		err = os.MkdirAll(parent, 0755)

		if err != nil {
			fatal(err)
		}

		filename := generator.ProxyPackage + ".go"

		path := filepath.Join(parent, filename)

		err := ioutil.WriteFile(path, code, 0644)

		if err != nil {
			fatal(err)
		}
	} else {
		fmt.Println(string(code))
	}
}
