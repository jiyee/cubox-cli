package main

import (
	"github.com/jiyee/cubox-cli/client"
	"github.com/pkg/errors"
	"log"
	"os"
)

type LinkCommand struct {
	Verbose bool `short:"V" long:"verbose" description:"show verbose debug information"`

	API string `long:"api" description:"cubox custom API" env:"CUBOX_API"`

	Content string `long:"url" description:"the url of the link"`

	Title string `short:"t" long:"title" description:"the title of the link"`

	Description string `short:"d" long:"desc" description:"the description of the link"`

	Tags []string `long:"tag" description:"additional tags"`

	Folder string `short:"f" long:"folder" description:"specify a folder for the url"`
}

func (x *LinkCommand) Usage() string {
	return "[options] <the link of url>"
}

func (x *LinkCommand) Execute(args []string) error {
	code, err := x.Run(args)

	os.Exit(code)

	return err
}

func (x *LinkCommand) Run(args []string) (int, error) {
	if x.Content == "" {
		return 1, errors.New("must specify the url of link")
	}

	link := client.Link{
		Type:        "url",
		Content:     x.Content,
		Title:       x.Title,
		Description: x.Description,
		Tags:        x.Tags,
		Folder:      x.Folder,
		API:         x.API,
	}

	responseMessage, err := link.Submit(x.Verbose)

	if err != nil {
		log.Println(err)
		return 1, err
	}

	log.Println(*responseMessage)
	return 0, nil
}
