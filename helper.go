package main

import (
	jira "github.com/andygrunwald/go-jira"
)

func (c *Client) findIssues(jql, issueKey string) (*[]jira.Issue, error) {
	if len(jql) > 0 {
		issues, err := c.Search(jql)

		if err != nil {
			return nil, err
		}

		return issues, nil
	} else if len(issueKey) > 0 {
		issue, err := c.Issue(issueKey)

		if err != nil {
			return nil, err
		}

		issues := []jira.Issue{*issue}
		return &issues, nil
	} else {
		panic("jql or issueKey is required.")
	}
}
