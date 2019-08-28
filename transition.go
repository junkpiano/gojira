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

			if len(jql) > 0 {
				issues, err := jc.Search(jql)

				if err != nil {
					panic(err)
				}

				if list {
					for _, issue := range *issues {
						ts, _ := jc.Transitions(issue.Key)
						for _, t := range *ts {
							fmt.Printf("%s: %s\n", t.Name, t.ID)
						}
					}
					return nil
				} else if len(action) == 0 {
					fmt.Printf("Action is required.\n\n")
					return nil
				}

				for _, issue := range *issues {
					err := jc.DoTransition(issue.Key, c.String("action"))

					if err != nil {
						panic(err)
					}
				}
			} else if len(issueKey) > 0 {
				issue, err := jc.Issue(issueKey)

				if err != nil {
					panic(err)
				}

				if list {
					ts, _ := jc.Transitions(issue.Key)
					for _, t := range *ts {
						fmt.Printf("%s: %s\n", t.Name, t.ID)
					}
					return nil
				} else if len(action) == 0 {
					fmt.Printf("Action is required.\n\n")
					return nil
				}

				err = jc.DoTransition(issue.Key, c.String("action"))

				if err != nil {
					panic(err)
				}
			} else {
				fmt.Printf("Either jql or issue is required.")
				return nil
			}

			return nil
		},
	}

	return command
}

func listAction() {

}
