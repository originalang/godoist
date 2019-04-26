# Togoist

> A simple package for interacting with the todoist Sync API written in Go.

## Installation

**Note:** Ensure that the Go [workspace](https://golang.org/doc/code.html#Workspaces) is set up before proceeding.

Run this command:

```
$ go get github.com/originalang/togoist
```
This will download the repository in your workspace and allow you to import it in your projects.

## Quick Start

All interaction with the todoist Sync API is managed through a client. A todoist API token is required to initialize the client. Navigate to [this](https://todoist.com/prefs/integrations) page and follow the instructions to issue a new token. 

With your new token, do the following:

```go
package main

import (
  "github.com/originalang/togoist"
)

func main() {
  // replace this with your personal token
  token := "YOUR_API_TOKEN"

  // initialize the client
  client := togoist.NewClient(token)

  // use the sync method to retrieve user, project, and item information
  client.Sync()
}
```

After this initial setup, you can refer to fields on the client struct to retrieve key information:

```go
// loop through all projects
for _, proj := range client.Projects {
  fmt.Println(proj.Name)
}

// loop through items
for _, itm := range client.Items {
  fmt.Printf("%v -- %s", itm.Id, itm.Content)
}
```

The following fields are available on the ```Project``` struct:

```go
Id         int64
Name       string
Color      int
Indent     int
Order      int
Collapsed  int
Shared     bool
IsDeleted  int
IsArchived int
IsFavorite int
```

The following fields are available on the ```Item``` struct:

```go
AllDay            bool
AssignedBy        int64
Checked           int
Collapsed         int
Content           string
DateAdded         string
DateCompleted     string
DateString        string
DayOrder          int
DueDate           string
Id                int64
InHistory         int
Indent            int
IsArchived        int
IsDeleted         int
ItemOrder         int
ParentId          int64
Priority          int
ProjectId         int64
ResponsibleUserId int64
SyncId            string
UserId            int64
```

Use these functions to interact with the todoist API:

```go
// methods
client.AddProject(name string, indent int)
client.UpdateProject(p Project)
client.DeleteProjects(ids []int64)
client.ArchiveProjects(ids []int64)
client.UnarchiveProjects(ids []int64)
client.AddItem(projectId int64, content string, indent int, dueDate string)
client.UpdateItem(item Item)
client.DeleteItems(ids []int64)
client.CompleteItems(ids []int64, toHistory bool)
client.UncompleteItems(ids []int64)

// functions
GetProjectId(c *Client, name string)
GetChildrenIds(c *Client, parentId int64)
```