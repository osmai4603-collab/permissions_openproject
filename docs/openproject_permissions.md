# OpenProject Permissions Guide

> **Source**: [OpenProject Permissions Guide](https://www.openproject.org/docs/system-admin-guide/users-permissions/permissions-guide/)
>
> This document catalogs every permission available in OpenProject, organized by module.

---

## Table of Contents

1. [Project](#1-project)
2. [Work Packages & Gantt](#2-work-packages--gantt)
3. [Boards](#3-boards)
4. [Backlogs (Scrum)](#4-backlogs-scrum)
5. [Budgets](#5-budgets)
6. [Calendars](#6-calendars)
7. [Documents](#7-documents)
8. [Forums](#8-forums)
9. [GitHub Integration](#9-github-integration)
10. [GitLab Integration](#10-gitlab-integration)
11. [Meetings](#11-meetings)
12. [News](#12-news)
13. [Team Planner](#13-team-planner)
14. [Time & Costs](#14-time--costs)
15. [Wiki](#15-wiki)
16. [Common Permission Dependencies](#16-common-permission-dependencies)

---

## 1. Project

| Permission | Description |
| --- | --- |
| Archive project | Allows archiving the project. |
| Create subproject | Allows creating subprojects under this project. |
| Edit project | Allows editing the project name, description, and settings. |
| Copy projects | Allows copying the project (requires Edit project). |
| Select project modules | Allows enabling or disabling project modules. |
| Manage members | Allows adding, removing, and changing roles of project members. |
| Invite members by email | Allows inviting new members via email (requires Manage members). |
| Manage versions | Allows creating, editing, and deleting project versions. |
| Manage types | Allows selecting which work package types are available. |
| Manage project custom fields | Allows managing project-level custom fields. |
| Add work packages | Allows creating new work packages. |
| View work packages | Allows viewing work packages. |
| Edit work packages | Allows editing work packages. |
| Move work packages | Allows moving work packages to other projects. |
| Copy work packages | Allows copying work packages. |
| Delete work packages | Allows deleting work packages. |
| Manage work package relations | Allows managing relations between work packages. |
| Manage subtasks | Allows creating and managing subtasks. |
| Add notes | Allows adding comments/notes to work packages. |
| View own work packages | Allows viewing only own work packages. |
| View work packages in calendar | Allows viewing work packages in calendar view. |
| Manage public queries | Allows managing shared/public work package queries. |
| Save queries | Allows saving personal work package queries. |
| Share work packages | Allows sharing work packages with non-members. |
| Automatically managed project folders — Share files | Allows sharing files in auto-managed project folders. |

## 2. Work Packages & Gantt

| Permission | Description |
| --- | --- |
| View work packages | Allows viewing work packages. |
| Add work packages | Allows creating new work packages in the project. |
| Edit work packages | Allows editing existing work packages. |
| Add attachments | Allows adding attachments to work packages (independent of Edit). |
| Delete work packages | Allows deleting work packages. |
| Move work packages | Allows moving work packages between projects. |
| Copy work packages | Allows copying work packages. |
| Add notes | Allows adding comments to work packages. |
| View own work packages | Allows viewing only own work packages. |
| Manage work package relations | Allows managing relations (blocks, follows, etc.). |
| Manage subtasks | Allows creating and managing subtask hierarchies. |
| Manage public queries | Allows managing shared/saved queries. |
| Save queries | Allows saving personal queries. |
| Assign versions | Allows assigning work packages to versions. |
| Change work package status | Allows changing work package status (independent of Edit). |
| Become assignee or responsible | Controls if user appears in assignee/responsible dropdowns. |
| Share work packages | Allows sharing work packages with non-members. |
| Manage work package views (Gantt) | Allows creating and managing Gantt chart views. |
| View file links | Allows viewing file links on work packages. |
| Manage file links | Allows creating and deleting file links on work packages. |
| Manage wiki page links | Allows linking wiki pages to work packages. |
| Create baseline | Allows creating baselines for comparing work package state over time. |
| Show baseline | Allows viewing existing baselines. |
| Export work packages | Allows exporting work packages to CSV, PDF, or Atom. |
| Select custom fields | Allows selecting which custom fields appear on work packages. |
| View spent time | Allows viewing tracked time entries. |
| Comment on work packages | Allows commenting on work packages. |
| Create work packages via email | Allows creating work packages by sending email. |
| Add watchers | Allows adding watchers to work packages. |
| Delete watchers | Allows removing watchers from work packages. |
| View commit messages | Allows viewing commit messages linked to work packages. |

## 3. Boards

| Permission | Description |
| --- | --- |
| View boards | Allows viewing Kanban-style boards. |
| Manage boards | Allows creating, editing, and deleting boards. |

## 4. Backlogs (Scrum)

| Permission | Description |
| --- | --- |
| View sprints | Allows viewing sprints in the backlogs module. |
| Manage sprints | Allows creating, editing, and deleting sprints. |
| View backlogs | Allows viewing the product backlog. |
| View task boards | Allows viewing task boards for sprints. |
| Manage task boards | Allows creating, editing, and deleting task boards. |
| Share sprint | Allows sharing sprint details with others. |

## 5. Budgets

| Permission | Description |
| --- | --- |
| View budgets | Allows viewing budgets within the project. |
| Edit budgets | Allows creating, editing, and deleting budgets. |

## 6. Calendars

| Permission | Description |
| --- | --- |
| View calendars | Allows viewing calendars. |
| Edit calendars | Allows creating, editing, and deleting calendars. |
| Subscribe to iCalendars | Allows subscribing to external calendar feeds. |

## 7. Documents

| Permission | Description |
| --- | --- |
| View documents | Allows viewing documents in the project. |
| Manage documents | Allows uploading, editing, and deleting documents. |

## 8. Forums

| Permission | Description |
| --- | --- |
| Manage forums | Allows creating, editing, and deleting forums. |
| Post messages | Allows posting messages in forums. |
| Edit messages | Allows editing any forum message. |
| Edit own messages | Allows editing only own forum messages. |
| Delete messages | Allows deleting any forum message. |
| Delete own messages | Allows deleting only own forum messages. |

## 9. GitHub Integration

| Permission | Description |
| --- | --- |
| Show GitHub content | Allows viewing GitHub pull requests and issues linked to work packages. |

## 10. GitLab Integration

| Permission | Description |
| --- | --- |
| Show GitLab content | Allows viewing GitLab merge requests and issues linked to work packages. |

## 11. Meetings

| Permission | Description |
| --- | --- |
| View meetings | Allows viewing meetings in a project. |
| Create meetings | Allows creating new meetings. |
| Edit meetings | Allows editing existing meetings. |
| Delete meetings | Allows deleting meetings. |
| Send meeting invites | Allows sending email invitations to meeting participants. |
| Manage agendas | Allows creating, editing, and closing meeting agendas. |
| Manage outcomes | Allows creating, editing, and closing meeting outcomes/minutes. |

## 12. News

| Permission | Description |
| --- | --- |
| Manage news | Allows creating, editing, and deleting news entries. |
| Comment news | Allows posting comments on news entries. |

## 13. Team Planner

| Permission | Description |
| --- | --- |
| View team planner | Allows viewing team planner views. |
| Manage team planner | Allows creating, editing, and deleting team planner views. |

## 14. Time & Costs

| Permission | Description |
| --- | --- |
| View spent time | Allows viewing time entries logged by all users. |
| View own spent time | Allows viewing only own time entries. |
| Log own time | Allows logging time on work packages for oneself. |
| Log time for other users | Allows logging time entries on behalf of other users. |
| Edit own time logs | Allows editing own logged time entries. |
| Edit time logs for other users | Allows editing time entries logged by other users. |
| Manage project activities | Allows managing time tracking activity types. |
| View own hourly rate | Allows viewing own hourly labor rate. |
| View all hourly rates | Allows viewing hourly rates of all members. |
| Edit own hourly rates | Allows editing own hourly labor rate. |
| Edit hourly rates | Allows editing hourly rates for all members. |
| View cost rates | Allows viewing cost types and their rates. |
| Book unit costs for oneself | Allows booking unit costs for oneself. |
| Book unit costs | Allows booking unit costs for all members. |
| Edit own booked unit costs | Allows editing own booked unit cost entries. |
| Edit booked unit costs | Allows editing unit cost entries booked by any member. |
| View booked costs | Allows viewing booked costs for all members. |
| View own booked costs | Allows viewing only own booked costs. |
| Save public cost reports | Allows saving cost reports visible to all members. |
| Save private cost reports | Allows saving cost reports visible only to oneself. |

## 15. Wiki

| Permission | Description |
| --- | --- |
| View wiki | Allows viewing wiki pages. |
| View wiki history | Allows viewing the edit history of wiki pages. |
| Edit wiki pages | Allows creating and editing wiki pages. |
| Manage wiki | Allows managing the wiki including renaming, deleting pages, and menu. |

---

## 16. Common Permission Dependencies

| Permission | Dependency | Notes |
| --- | --- | --- |
| Invite members by email | Manage members | Cannot invite without member management access. |
| Copy projects | Edit project | Copying requires project editing access. |
| Add attachments | *(independent)* | Does not require Edit work packages. |
| Change work package status | *(independent)* | Does not require Edit work packages. |
| Become assignee or responsible | *(behavioral)* | Controls whether the user appears in assignee/responsible selection dropdowns. |

---

## Statistics

| Metric | Count |
| --- | --- |
| **Total modules** | 15 |
| **Total permissions** | ~100 |
| **Project-level permissions** | ~100 |
| **Global-level permissions** | 0 (all are project-scoped) |

---

*Generated from the official [OpenProject Permissions Guide](https://www.openproject.org/docs/system-admin-guide/users-permissions/permissions-guide/).*
