package main

import (
	"log"
	_ "net/http/pprof"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var markdown bool

func main() {
	app := cli.NewApp()
	app.Name = "Frank"
	app.Version = "20171221"
	app.Usage = "Command line REST API automated testing tool"
	app.Authors = []*cli.Author{
		{
			Name:  "Cloud",
			Email: "cloud@txthinking.com",
		},
	}
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "markdown, m",
			Usage:       "Print markdown document",
			Destination: &markdown,
		},
		&cli.Int64Flag{
			Name:  "delay, d",
			Usage: "Delay per request, (millisecond)",
			Value: 0,
		},
		&cli.StringFlag{
			Name:  "case, c",
			Usage: "Path of case file",
			Value: "case.frank",
		},
	}
	app.Action = func(c *cli.Context) error {
		if err := runCase(c.String("case"), c.Int64("delay")); err != nil {
			color.Red(err.Error())
		}
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func runCase(f string, d int64) error {
	c, err := NewCase(f, d)
	if err != nil {
		return err
	}
	return c.Run()
}
