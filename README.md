# j
The jira ticket issuer

## installation

```shell
$ go get github.com/roronya/j
```


## usage

```shell
$ export JIRA_USER="JIRA_USER"
$ export JIRA_PASSWORD="JIRA_PASSWORD"
$ export JIRA_SERVER="https://your.jira.server.com/"
$ j -p PROJECT -c COMPONENT -e EPIC -s SUMMARY -d DESCRIPTION
```
