package main

import (
	"flag"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"log"
	"os"
)

var project string
var epic string
var component string
var issueType string
var summary string
var description string
var assigneeEmail string
var reporterEmail string

// main
// example: j -p MYPRJ -s "something to do"
func main() {
	flag.StringVar(&project, "p", "", "[must] Project")
	flag.StringVar(&component, "c", "", "[option] Component")
	flag.StringVar(&epic, "e", "", "[option] Epic")
	flag.StringVar(&issueType, "i", "Task", "[option] Issue Type") // TODO: EPICに対応する
	flag.StringVar(&summary, "s", "", "[must] Summary")
	flag.StringVar(&description, "d", "", "[option] Description")
	flag.StringVar(&assigneeEmail, "a", "", "[option] Assignee Email Address")
	flag.StringVar(&reporterEmail, "r", "", "[option] Reporter Email Address")
	flag.Parse()
	if project == "" || summary == "" {
		flag.Usage()
		return
	}

	user := os.Getenv("JIRA_USER")
	pass := os.Getenv("JIRA_PASSWORD")
	server := os.Getenv("JIRA_SERVER")
	client, err := NewClient(user, pass, server)
	if err != nil {
		log.Fatal(err)
	}

	var assigneeId string
	if assigneeEmail != "" {
		assignee, err := GetUser(client, assigneeEmail)
		if err != nil {
			log.Fatal(err)
		}
		assigneeId = assignee.AccountID
	}

	var reporterId string
	if reporterEmail != "" {
		reporter, err := GetUser(client, reporterEmail)
		if err != nil {
			log.Fatal(err)
		}
		reporterId = reporter.AccountID
	}

	url, err := IssueCreate(
		client,
		project,
		component,
		epic,
		issueType,
		summary,
		assigneeId,
		reporterId,
		description,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(url)
}

func NewClient(
	username string,
	password string,
	baseUrl string,
) (*jira.Client, error) {
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}
	client, err := jira.NewClient(tp.Client(), baseUrl)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetUser
// emailをもとにUserオブジェクトを探して返す
// emailに紐づくユーザーが見つからなかった場合はerrorを返す
func GetUser(
	client *jira.Client,
	email string,
) (*jira.User, error) {
	users, _, err := client.User.Find(email)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("couldn't find a such user<%s>", email)
	}
	return &users[0], nil
}

// IssueCreate
// client: NewClientによって作ったclient
// project: JIRAのプロジェクト。存在するJIRAのプロジェクトを必ず渡す。存在しないJIRAのプロジェクトを渡すとエラーになる。
// component: コンポーネント。設定しない場合は空文字を渡す。
// epic: エピック。設定しない場合は空文字を渡す。
// issueType: 現状はTaskのみ対応している。Task以外を渡すとエラーになる。
// summary: JIRAのチケットのタイトルになる
// assignee: アサインする人のアカウントIDを渡す。アカウントIDはGetUserによって取得できる。アサインしない場合は空文字を渡すと非アサイン状態になる。
// reporter: アサインする人のアカウントIDを渡す。アカウントIDはGetUserによって取得できる。アサインしない場合は空文字を渡すとチケットを作った人が報告者になる。
// description: 説明文
func IssueCreate(
	client *jira.Client,
	project string,
	component string,
	epic string,
	issueType string,
	summary string,
	assignee string,
	reporter string,
	description string,
) (url string, err error) {
	f := &jira.IssueFields{
		Project: jira.Project{
			Key: project,
		},
		Type: jira.IssueType{
			Name: issueType,
		},
		Summary:     summary,
		Description: description,
	}
	if component != "" {
		f.Components = []*jira.Component{{Name: component}}
	}
	// see: https://github.com/andygrunwald/go-jira/issues/307
	// FIXME: epicはcustomefieldとして作られていて、プロジェクトによって違う
	if epic != "" {
		f.Unknowns = map[string]interface{}{
			"customfield_10006": epic,
		}
	}
	if assignee != "" {
		f.Assignee = &jira.User{AccountID: assignee}
	}
	if reporter != "" {
		f.Reporter = &jira.User{AccountID: reporter}
	}
	i := jira.Issue{Fields: f}
	issue, _, err := client.Issue.Create(&i)
	if err != nil {
		return "", err
	}
	baseURL := client.GetBaseURL()
	return fmt.Sprintf("%sbrowse/%s", baseURL.String(), issue.Key), nil
}
