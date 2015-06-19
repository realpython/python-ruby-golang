package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gomr"
	app.Author = "Kyle W. Purdon"
	app.Version = "0.0.1"
	app.Usage = "multi-repository manager in go"
	app.Commands = []cli.Command{
		{
			Name:   "register",
			Usage:  "register a directory",
			Action: register,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "directory, d",
					Value: "./",
					Usage: "directory to tag",
				},
				cli.StringFlag{
					Name:  "tag, t",
					Value: "default",
					Usage: "tag to add for directory",
				},
				cli.BoolFlag{
					Name:  "append",
					Usage: "append the tag to an existing registered directory",
				},
			},
		},
		{
			Name:   "run",
			Usage:  "run a command",
			Action: run,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "basepath, b",
					Value: "./",
					Usage: "path to begin the recursive search",
				},
				cli.StringFlag{
					Name:  "tag, t",
					Value: "default",
					Usage: "tag to add for directory",
				},
				cli.BoolFlag{
					Name:  "dryrun, d",
					Usage: "print (dont execute) the commands that will be run",
				},
			},
		},
	}

	app.Run(os.Args)
}
