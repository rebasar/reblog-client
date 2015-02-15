package main

import (
	"fmt"
	"github.com/ogier/pflag"
	"io"
	"os"
)

func usage() {
	fmt.Printf("Usage: %s [parameters]:\n\n", os.Args[0])
	pflag.PrintDefaults()
	fmt.Println()
}

func openInputFile() io.Reader {
	f, err := os.Open(InputFile)
	checkError(err)
	return f
}

func getInputFile() io.Reader {
	if InputFile == "-" {
		return os.Stdin
	} else {
		return openInputFile()
	}
}

func main() {
	initConfig()
	if Help {
		usage()
	} else {
		entry := readEntry(getInputFile())
		uploadEntry(entry)
	}
}
