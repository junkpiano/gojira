package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func TransitionCommand() cli.Command {
	command := cli.Command{
		Name:    "transition",
		Aliases: []string{"t"},
		Usage:   "perform transition",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "issue, i"},
			cli.StringFlag{Name: "action, a"},
			cli.StringFlag{Name: "jql, j"},
			cli.BoolFlag{Name: "list, l"},
			cli.StringFlag{Name: "comment, c"},
		},
		Action: func(c *cli.Context) error {
			jc, err := NewClient()

			if err != nil {
				panic(err)
			}

			action := c.String("action")
			jql := c.String("jql")
			issueKey := c.String("issue")
			list := c.Bool("list")
			comment := c.String("comment")

			issues, err := jc.findIssues(jql, issueKey)

			if err != nil {
				panic(err)
			}

			if list {
				for _, issue := range *issues {
					ts, _ := jc.Transitions(issue.Key)
					for _, t := range *ts {
						fmt.Printf("* %s: %s\n", t.Name, t.ID)
					}
				}
				return nil
			} else if len(action) == 0 {
				fmt.Printf("Action is required.\n\n")
				return nil
			}

			for _, issue := range *issues {
				if len(comment) == 0 {
					err = jc.DoTransition(issue.Key, action)
				} else {
					err = jc.DoTransitionWithPayload(issue.Key, action, comment)
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
