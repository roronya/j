# j
jira ticket issuer

## installation

```shell
$ brew install python3
$ pip3 install git+https://github.com/roronya/j
```

## usage

```shell
$ export JIRA_USER="JIRA_USER"
$ export JIRA_PASSWORD="JIRA_PASSWORD"
$ export JIRA_SERVER="https://your.jira.server.com"
$ j project epic component assignee title -d description
```

## develop

```shell
$ vim j/__main__.py
$ python3 -m j
```