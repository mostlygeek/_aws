package _ec2

import "github.com/codegangsta/cli"

func AddToCLI(a *cli.App) {
	commands := []cli.Command{
		{
			Name:    "ec2",
			Usage:   "find EC2 instances that match a regex",
			Aliases: []string{"e"},
			Subcommands: []cli.Command{
				{
					Name:    "match",
					Aliases: []string{"m"},
					Usage:   "find EC2 instances that match a regex",
					Action:  Match,
				},
			},
		},
	}

	a.Commands = append(a.Commands, commands...)
}
