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

			jql := c.String("jql")
			issueKey := c.String("issue")
			issues, err := jc.findIssues(jql, issueKey)

			payload := c.String("payload")

			if len(payload) == 0 {
				fmt.Printf("payload json is required.")
				return nil
			}

			var data map[string]interface{}
			err = json.Unmarshal([]byte(payload), &data)
			if err != nil {
				panic(err)
			}

			for _, issue := range *issues {
				err = jc.UpdateIssue(issue.Key, data)
				if err != nil {
					panic(err)
				}
			}

			return nil
		},
	}

	return command
}
