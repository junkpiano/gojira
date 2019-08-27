package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		IssuesCommand(),
		TransitionCommand(),
		AssigneeCommand(),
		UpdateIssueCommand(),
	}

	app.Run(os.Args)
}
