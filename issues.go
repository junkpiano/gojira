package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func IssuesCommand() cli.Command {
	command := cli.Command{
		Name:    "issues",
		Aliases: []string{"i"},
		Usage:   "list issues",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "jql, j"},
			cli.StringFlag{Name: "issue, i"},
		},
		Action: func(c *cli.Context) error {
			jc, err := NewClient()

			if err != nil {
				panic(err)
			}

			jql := c.String("jql")
			issueKey := c.String("issue")
			issues, err := jc.findIssues(jql, issueKey)

			if err != nil {
				panic(err)
			}

			for _, issue := range *issues {
				fmt.Printf("* %s: %s\n -> %sbrowse/%s\n", issue.Key, issue.Fields.Summary, os.Getenv("GOJIRA_BASEURL"), issue.Key)
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
		Flags: []cli.Flag{
			cli.StringFlag{Name: "jql, j"},
			cli.StringFlag{Name: "issue, i"},
			cli.StringFlag{Name: "payload, p"},
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

			jql := c.String("jql")

			issues, err := jc.findIssues(jql, issueKey)

			if c.NArg() > 1 {
				payload := c.Args().Get(1)
				var data map[string]interface{}
				err := json.Unmarshal([]byte(payload), &data)
				if err != nil {
					panic(err)
				}

				for _, issue := range *issues {
					err = jc.UpdateIssue(issue.Key, data)
					if err != nil {
						panic(err)
					}
				}
			} else {
				fmt.Println("json is required")
			}

			return err
		},
	}

	return command
}
