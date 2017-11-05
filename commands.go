package main

import (
	"fmt"
	"os"

	"github.com/Ladicle/toggl/command"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name: "cache",
	},
	cli.BoolFlag{
		Name: "csv",
	},
}

var projectIDFlag = cli.IntFlag{
	Name:  "project-id, P",
	Usage: "Project id",
}

var Commands = []cli.Command{
	{
		Name:   "start",
		Usage:  "Start time entry",
		Action: command.CmdStart,
		Flags: []cli.Flag{
			projectIDFlag,
		},
	},
	{
		Name:   "stop",
		Usage:  "End time entry",
		Action: command.CmdStop,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "current",
		Usage:  "Show current time entry",
		Action: command.CmdCurrent,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "workspaces",
		Usage:  "Show workspaces",
		Action: command.CmdWorkspaces,
	},
	{
		Name:   "projects",
		Usage:  "Show projects on current workspaces",
		Subcommands: []cli.Command{
			{
				Name:   "show",
				Usage:  "Show current workspace projects",
				Action: command.CmdShowProjects,
			},
			{
				Name:   "create",
				Usage:  "Create current workspace project",
				Action: command.CmdCreateProject,
			},
			{
				Name:   "delete",
				Usage:  "Delete project",
				Action: command.CmdDeleteProject,
			},
		},
	},
	{
		Name:   "local",
		Usage:  "Set current dir workspace",
		Action: CmdLocal,
	},
	{
		Name:   "global",
		Usage:  "Set global workspace",
		Action: CmdGlobal,
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
