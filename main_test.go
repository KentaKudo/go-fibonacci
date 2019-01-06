package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestStatusCode(t *testing.T) {
	sut := &CLI{os.Stdout, os.Stderr}
	args := strings.Split("-n 10", " ")
	if status := sut.Run(args); status != ExitCodeOK {
		t.Errorf("got %d, want %d", status, ExitCodeOK)
	}
}

func TestRun(t *testing.T) {
	outSteam, errStream := new(bytes.Buffer), new(bytes.Buffer)
	sut := &CLI{outStream: outSteam, errStream: errStream}
	args := strings.Split("-n 10", " ")

	sut.Run(args)
	want := fmt.Sprintf("fib(%d) = %d\n", 10, 55)
	if !strings.Contains(outSteam.String(), want) {
		t.Errorf("got %q, want %q", outSteam.String(), want)
	}
}
