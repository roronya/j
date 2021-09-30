from jira import JIRA
import os

USER = os.environ['JIRA_USER']
PASSWORD = os.environ['JIRA_PASSWORD']
SERVER = os.environ['JIRA_SERVER']


def main(project, summary, epic, component, assignee, description, issuetype, parent):
    options = {'server': SERVER}
    j = JIRA(options=options, basic_auth=(USER, PASSWORD))
    issue_dict = {'project': project, 'summary': summary,
                  'priority': {'id': '3'},
                  'issuetype': {'name': issuetype if issuetype is not None else 'Task'}}
    if description:
        issue_dict['description'] = description
    if assignee:
        issue_dict['assignee'] = {'name': assignee}
    if component:
        issue_dict['components'] = [{"name": component}]
    new_issue = j.create_issue(issue_dict)
    if epic:
        j.add_issues_to_epic(epic, new_issue.key)
    return f"{SERVER}/browse/{new_issue.key}"
