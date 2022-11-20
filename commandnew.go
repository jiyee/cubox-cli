package main

import (
	"bufio"
	"fmt"
	"github.com/jiyee/cubox-cli/client"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"strings"
)

type NewCommand struct {
	Verbose bool `short:"V" long:"verbose" description:"show verbose debug information"`

	API string `long:"api" description:"cubox custom API" env:"CUBOX_API"`

	Tags []string `short:"t" long:"tag" description:"additional tags"`

	Folder string `short:"f" long:"folder" description:"specify a folder for the memo"`
}

func (x *NewCommand) Usage() string {
	return "[new command options] <the content of memo>"
}

func (x *NewCommand) Execute(args []string) error {
	code, err := x.Run(args)

	os.Exit(code)

	return err
}

func (x *NewCommand) Run(args []string) (int, error) {
	var content string

	if isInputFromPipe() {
		content = getStdinContent(os.Stdin)
	} else {
		content = strings.Join(args, " ")
	}

	if content == "" {
		return 1, errors.New("must specify the content of memo")
	}

	memo := client.Memo{
		Type:    "memo",
		Content: content,
		Tags:    x.Tags,
		Folder:  x.Folder,
		API:     x.API,
	}

	responseMessage, err := memo.Submit(x.Verbose)

	if err != nil {
		log.Println(err)
		return 1, err
	}

	log.Println(*responseMessage)
	return 0, nil
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func getStdinContent(r io.Reader) string {
	var runes []rune
	var output string

	reader := bufio.NewReader(r)

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		runes = append(runes, input)
	}

	for j := 0; j < len(runes); j++ {
		output += fmt.Sprintf("%c", runes[j])
	}

	return output
}
