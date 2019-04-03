/*
 * Command line tool for pull request form Gitlab repos
 */
package main

import (
	"github.com/urfave/cli"
	"github.com/wuzuoliang/gitcli/cmd"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

const (
	AppName = "gitcli"
)

// global config
var conf = struct {
	Verbose bool
}{}

func main() {
	app := cli.NewApp()
	app.Name = AppName
	app.Version = "0.0.1"
	app.Usage = "Command line tool to pull GitLab repos. Support on Windows and Mac OS"

	// global flag
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "verbose",
			Usage:       "enable verbose mode to dump debug info",
			Destination: &conf.Verbose,
		},
	}

	// sub commands
	app.Commands = []cli.Command{
		{
			Name:   "configure",
			Usage:  "Configure this app",
			Action: cmd.Configure,
		},
		{
			Name:  "get",
			Usage: "Pull all directory under the path",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "output-dir, o",
					Usage:       "Specify the output `DIRECTORY`",
					Destination: &cmd.GetParams.OutputDirectory,
				},
				cli.StringSliceFlag{
					Name:  "exclude, e",
					Usage: "Exclude specified `DIRECTORY",
					Value: &cmd.GetParams.ExcludeList,
				},
			},
			Action: cmd.GetHandler,
		},
		{
			Name:   "tree",
			Usage:  "tree all repo or project of current user",
			Action: cmd.TreeHandler,
		},
		{
			Name:   "ls",
			Usage:  "ls all repo or project of current user",
			Action: cmd.LsHandler,
		},
		{
			Name:   "reset",
			Usage:  "Remove user configuration of this app",
			Action: cmd.RemoveConfig,
		},
		{
			Name:   "load",
			Usage:  "load git directory into local db",
			Action: cmd.LoadHandler,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	// handle global options
	app.Before = func(c *cli.Context) error {
		// set log output
		if conf.Verbose {
			log.SetOutput(os.Stderr)
		} else {
			log.SetOutput(ioutil.Discard)
		}
		return nil
	}

	app.Action = cli.ShowAppHelp

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
