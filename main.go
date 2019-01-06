package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
	ExitCodeFileError
)

type CLI struct {
	outStream, errStream io.Writer
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args[1:]))
}

func (c *CLI) Run(args []string) int {
	flagset := flag.NewFlagSet("fibonacci", flag.ExitOnError)
	var n uint
	flagset.UintVar(&n, "n", 0, "The digit to get the fibbonachi number")
	flagset.Parse(args)

	fmt.Fprintf(c.outStream, "fib(%d) = %d\n", n, fibonacci(n))
	return 0
}
