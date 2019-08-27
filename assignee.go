package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func AssigneeCommand() cli.Command {
	command := cli.Command{
		Name:    "assignee",
		Aliases: []string{"a"},
		Usage:   "change assignee",
		Flags: []cli.Flag{
			cli.BoolFlag{Name: "reporter, r"},
		},
		Action: func(c *cli.Context) error {
			jc, err := NewClient()

			if err != nil {
				panic(err)
			}

			var issueKey string

			if c.NArg() > 0 {
				issueKey = c.Args().Get(0)
			} else {
				fmt.Println("IssueKey is required.")
				return nil
			}

			if c.NArg() > 1 {
				userName := c.Args().Get(1)
				jc.UpdateAssignee(issueKey, &userName)
			} else {
				if c.Bool("reporter") {
					issue, err := jc.Issue(issueKey)

					if err != nil {
						panic(err)
					}

					jc.UpdateAssignee(issueKey, &(issue.Fields.Creator.Name))
				} else {
					jc.UpdateAssignee(issueKey, nil)
				}
			}

			return err
		},
	}

	return command
}
