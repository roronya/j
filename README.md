# j
jira ticket issuer

## installation

```shelｇl
$ brew install python3
$ pip3 install git+https://github.com/roronya/j.git # for mac
```

## usage

```shell
$ export JIRA_USER="JIRA_USER"
$ export JIRA_PASSWORD="JIRA_PASSWORD"
$ export JIRA_SERVER="https://your.jira.server.com"
$ j PROJECT summary # make a vanilla task by default
$ j PROJECT summary -e epic -c component # make a task assigning an epic and a component
$ j PROJECT summary -i story # make a story
$ j PROJECT summary -i Sub-task -p parent-task-key # make a sub-task
```

## develop

```shell
$ poetry install # setup
$ poetry run j # execute j/j.py in virtual env
```
