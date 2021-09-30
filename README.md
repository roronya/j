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
$ j PROJECT summary # デフォルトではバニラのタスクができる
$ j PROJECT summary -e epic -c component # エピックとコンポーネントを指定してタスクを作る
$ j PROJECT summary -i ストーリー # ストーリーを作る
$ j PROJECT summary -i Sub-task -p parent-task-key# サブタスクを作る
```

## develop

```shell
$ vim j/j.py
$ ./bin/j
```