package main

import (
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	var options struct{}
	var parser = flags.NewParser(&options, flags.Default)

	if _, err := parser.AddCommand("memo", "Create a new memo", "", &MemoCommand{}); err != nil {
		log.Fatal(err)
	}
	if _, err := parser.AddCommand("link", "Create a new link", "", &LinkCommand{}); err != nil {
		log.Fatal(err)
	}
	if _, err := parser.AddCommand("version", "Print version", "", &VersionCommand{}); err != nil {
		log.Fatal(err)
	}

	if _, err := parser.Parse(); err != nil {
		switch err.(type) {
		case *flags.Error:
			fe, _ := err.(*flags.Error)
			if fe.Type == flags.ErrHelp {
				return 0
			}
			return 1
		default:
			return 1
		}
	}

	return 0
}
