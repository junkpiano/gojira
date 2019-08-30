package main

import (
	"github.com/urfave/cli"
)

func AssigneeCommand() cli.Command {
	command := cli.Command{
		Name:    "assignee",
		Aliases: []string{"a"},
		Usage:   "change assignee",
		Flags: []cli.Flag{
			cli.BoolFlag{Name: "reporter, r"},
			cli.StringFlag{Name: "jql, j"},
			cli.StringFlag{Name: "issue, i"},
			cli.StringFlag{Name: "user, u"},
		},
		Action: func(c *cli.Context) error {
			jc, err := NewClient()

			if err != nil {
				panic(err)
			}

			jql := c.String("jql")
			issueKey := c.String("issue")
			user := c.String("user")
			reporter := c.Bool("reporter")
			issues, err := jc.findIssues(jql, issueKey)

			for _, issue := range *issues {
				var err error
				if reporter {
					err = jc.UpdateAssignee(issue.Key, &(issue.Fields.Creator.Name))
				} else {
					err = jc.UpdateAssignee(issue.Key, &user)
				}

				if err != nil {
					panic(err)
				}
			}

			return nil
		},
	}

	return command
}
