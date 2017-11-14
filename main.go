package main

import (
	"log"
	_ "net/http/pprof"
	"os"

	"github.com/urfave/cli"
)

var markdown bool

func main() {
	app := cli.NewApp()
	app.Name = "Frank"
	app.Version = "20171114"
	app.Usage = "Command line REST API testing tool"
	app.Author = "Cloud"
	app.Email = "cloud@txthinking.com"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "markdown, m",
			Usage:       "Print markdown document",
			Destination: &markdown,
		},
		cli.StringFlag{
			Name:  "case, c",
			Usage: "Path of case file",
			Value: "case.frank",
		},
	}
	app.Action = func(c *cli.Context) error {
		return runCase(c.String("case"))
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func runCase(f string) error {
	c, err := NewCase(f)
	if err != nil {
		return err
	}
	return c.Run()
}
