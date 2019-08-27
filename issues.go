package main

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
)

func IssuesCommand() cli.Command {
	command := cli.Command{
		Name:    "issues",
		Aliases: []string{"i"},
		Usage:   "list issues",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "jql, j"},
		},
		Action: func(c *cli.Context) error {
			jc, err := NewClient()

			if err != nil {
				panic(err)
			}

			jql := c.String("jql")
			issues, err := jc.Search(jql)

			if err != nil {
				panic(err)
			}

			for _, issue := range *issues {
				fmt.Printf("%s, %s, %s, %s, %s\n", issue.Key, issue.Fields.Summary, issue.Fields.Type.Name, issue.Fields.Creator.Name, issue.Fields.Assignee.Name)
			}

			return err
		},
	}

	return command
}

func UpdateIssueCommand() cli.Command {
	command := cli.Command{
		Name:    "update",
		Aliases: []string{"u"},
		Usage:   "Update Issue",
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
				payload := c.Args().Get(1)
				var data map[string]interface{}
				err := json.Unmarshal([]byte(payload), &data)
				if err != nil {
					panic(err)
				}

				fmt.Println(data)

				err = jc.UpdateIssue(issueKey, data)
				if err != nil {
					panic(err)
				}
			} else {
				fmt.Println("json is required")
			}

			return err
		},
	}

	return command
}
