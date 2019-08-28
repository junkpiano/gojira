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

### Transition issues

#### List of transitions

```bash
gojira transition --jql "project=Awesome and status=Open" --list
```

#### Perform transition

```bash
gojira transition --jql "project=RAwesome and status=Open" --action 111 # 111 is ID of the next lane.
```

### Issue Assigning

#### Assign

```bash
gojira assignee Awesome-1234 username
```

#### Assign back to creator

```
gojira assignee --reporter Awesome-1234
```
