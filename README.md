Gojira
==================

## Usage

### Authentication

Register credential in the environmental variables.

```bash
export GOJIRA_BASEURL=http://yourjira.com/
export GOJIRA_USERNAME=xxx.yyyy
export GOJIRA_PASSWORD=qwertyasdfgh
```

That's it. gojira reads that information on the fly.

### Search issues

```bash
gojira issues --jql "project=Awesome and status=Open"
```

```bash
gojira issues --issue AWESOME-123
```

### Transition issues

#### List of transitions

```bash
gojira transition --jql "project=Awesome and status=Open" --list
```

```bash
gojira transition --jql "project=Awesome and status=Open" --list --comment text
```

#### Perform transition

```bash
gojira transition --jql "project=RAwesome and status=Open" --action 111 # 111 is ID of the next lane.
```

```bash
gojira transition --jql "project=RAwesome and status=Open" --action 111 --comment text # 111 is ID of the next lane.
```

### Issue Assigning

#### Assign

```bash
gojira assignee --issue Awesome-1234 --user username
```

```bash
gojira assignee --jql "project=RAwesome and status=Open" --user username
```

#### Assign back to the reporter

```
gojira assignee --issue Awesome-1234 --reporter 
```

```bash
gojira assignee --jql "project=RAwesome and status=Open" --reporter
```

### Update issues

#### Update

As of payload spec, please find your edition from [this page](https://developer.atlassian.com/server/jira/platform/rest-apis/).
This tool uses `Edit issue` API.

```bash
gojira update --issue AWESOME-123 --payload "<json string>"
```

```bash
gojira update --jql "project=RAwesome and status=Open" --payload "<json string>"
```