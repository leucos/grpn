package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	rad := flag.Bool("rad", true, "Use radians")
	deg := flag.Bool("deg", false, "Use radians")

	flag.Parse()

	if *rad && *deg {
		return fmt.Errorf("unable to use -rad and -deg at the same time")
	}

	e := &engine{
		stack:    &Stack{},
		previous: &Stack{},
	}
	return runUI(e)
}
