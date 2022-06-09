import argparse
from jira import JIRA
import os

USER = os.environ['JIRA_USER']
PASSWORD = os.environ['JIRA_PASSWORD']
SERVER = os.environ['JIRA_SERVER']


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('project', help='project')
    parser.add_argument('summary', help='summary')
    parser.add_argument('-e', '--epic', help='epic')
    parser.add_argument('-c', '--component', help='component')
    parser.add_argument('-a', '--assignee', help='assignee')
    parser.add_argument('-d', '--description', help='description')
    parser.add_argument('-i', '--issuetype', help='issuetype')
    parser.add_argument('-p', '--parent', help='parent')
    args = parser.parse_args()
    url = create(**args.__dict__)
    print(url)


def create(project, summary, epic, component, assignee, description, issuetype,
    parent):
    options = {'server': SERVER}
    j = JIRA(options=options, basic_auth=(USER, PASSWORD))
    issue_dict = {'project': project, 'summary': summary,
                  'priority': {'id': '3'},
                  'issuetype': {
                      'name': issuetype if issuetype is not None else 'Task'}}
    if description:
        issue_dict['description'] = description
    if assignee:
        issue_dict['assignee'] = {'name': assignee}
    if component:
        issue_dict['components'] = [{"name": component}]
    if parent:
        issue_dict['parent'] = {'id': parent}
    new_issue = j.create_issue(issue_dict)
    if epic:
        j.add_issues_to_epic(epic, [new_issue.key])
    return f"{SERVER}/browse/{new_issue.key}"
