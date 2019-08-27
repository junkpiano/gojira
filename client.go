package main

import (
	"os"

	jira "github.com/andygrunwald/go-jira"
)

type Client struct {
	base   string
	tp     jira.BasicAuthTransport
	engine *jira.Client
}

func NewClient() (*Client, error) {
	base := os.Getenv("GOJIRA_BASEURL")
	tp := jira.BasicAuthTransport{
		Username: os.Getenv("GOJIRA_USERNAME"),
		Password: os.Getenv("GOJIRA_PASSWORD"),
	}

	jiraClient, err := jira.NewClient(tp.Client(), base)

	if err != nil {
		return nil, err
	}

	c := &Client{
		base:   base,
		tp:     tp,
		engine: jiraClient,
	}

	return c, nil
}

func (c *Client) Search(jql string) (*[]jira.Issue, error) {
	issues, _, err := c.engine.Issue.Search(jql, nil)

	if err != nil {
		return nil, err
	}

	return &issues, nil
}

func (c *Client) Issue(issueKey string) (*jira.Issue, error) {
	issue, _, err := c.engine.Issue.Get(issueKey, nil)
	return issue, err
}

func (c *Client) UpdateIssue(issueKey string, data map[string]interface{}) error {
	_, err := c.engine.Issue.UpdateIssue(issueKey, data)
	return err
}

func (c *Client) Transitions(issueKey string) (*[]jira.Transition, error) {
	t, _, err := c.engine.Issue.GetTransitions(issueKey)
	return &t, err
}

func (c *Client) DoTransition(issueKey, transitionID string) error {
	_, err := c.engine.Issue.DoTransition(issueKey, transitionID)
	return err
}

func (c *Client) UpdateAssignee(issueKey string, userName *string) error {
	var err error
	if userName != nil {
		user := jira.User{Name: *userName}
		_, err = c.engine.Issue.UpdateAssignee(issueKey, &user)
	} else {
		_, err = c.engine.Issue.UpdateAssignee(issueKey, nil)
	}
	return err
}
