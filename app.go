package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.2.5"

	app.Commands = []cli.Command{
		IssuesCommand(),
		TransitionCommand(),
		AssigneeCommand(),
		UpdateIssueCommand(),
	}

	app.Run(os.Args)
}
