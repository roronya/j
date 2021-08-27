from jira import JIRA
import os


def main(project, epic, component, assignee, title, description):
  user = os.environ['JIRA_USER']
  password = os.environ['JIRA_PASSWORD']
  server = os.environ['JIRA_SERVER']
  options = {'server': server}
  j = JIRA(options=options, basic_auth=(user, password))
  new_issue = j.create_issue(
      project=project,
      summary=title,
      description=description or "",
      issuetype={'name': 'Task'},
      priority={'id': '3'},
      assignee={'name': assignee},
      components=[{"name": component}],
  )
  j.add_issues_to_epic(epic, [new_issue.key])
  return f"{server}/browse/{new_issue.key}"
