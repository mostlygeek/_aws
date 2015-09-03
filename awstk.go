package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mostlygeek/awstk/_ec2"
)

func main() {

	app := cli.NewApp()
	app.Name = "awstk"
	app.Usage = "random tools for aws stuff"

	_ec2.AddToCLI(app)

	app.Run(os.Args)
}
