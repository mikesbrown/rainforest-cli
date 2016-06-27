package main

import (
	"io"
	"os"

	"gopkg.in/urfave/cli.v2"
)

var apiToken string
var baseURL = "https://app.rainforestqa.com/api/1"
var out io.Writer = os.Stdout

func main() {

	app := cli.NewApp()
	app.Name = "Rainforest CLI"
	app.Usage = "Command line utility for Rainforest QA"

	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "tags",
			Usage: "Filter tests by tag",
		},
		cli.StringFlag{
			Name:  "token",
			Usage: "Rainforest API token",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Run your tests on Rainforest",
			Action: func(c *cli.Context) error {
				apiToken = c.String("token")
				createRun(c)
				return nil
			},
		},

		{
			Name:  "folders",
			Usage: "Retreive folders on Rainforest",
			Action: func(c *cli.Context) error {
				apiToken = c.String("token")
				fetchResource("Folders")
				return nil
			},
		},

		{
			Name:  "sites",
			Usage: "Retreive sites on Rainforest",
			Action: func(c *cli.Context) error {
				apiToken = c.String("token")
				fetchResource("Sites")
				return nil
			},
		},

		{
			Name:  "browsers",
			Usage: "Retreive sites on Rainforest",
			Action: func(c *cli.Context) error {
				apiToken = c.String("token")
				fetchResource("Browsers")
				return nil
			},
		},
	}
	app.Run(os.Args)
}
