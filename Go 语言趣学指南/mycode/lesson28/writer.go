package main

import (
	"fmt"
	"io"
	"os"
)

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintf(sw.w, s)
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{w: f}
	sw.writeln("Errors are values.\n")
	sw.writeln("Don’t just check errors, handle them gracefully.\n")
	sw.writeln("Don't panic.\n")
	sw.writeln("Make the zero value useful.\n")
	sw.writeln("The bigger the interface, the weaker the abstraction.\n")
	sw.writeln("interface{} says nothing.\n")
	sw.writeln("Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.\n")
	sw.writeln("Documentation is for users.\n")
	sw.writeln("A little copying is better than a little dependency.\n")
	sw.writeln("Clear is better than clever.\n")
	sw.writeln("Concurrency is not parallelism.\n")
	sw.writeln("Don’t communicate by sharing memory, share memory by communicating.\n")
	sw.writeln("Channels orchestrate; mutexes serialize.\n")

	return sw.err
}

func main() {
	err := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
