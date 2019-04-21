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

  // initialie the client
  client := togoist.NewClient(token)

  // use the sync method to retrieve user, project, and item information
  client.Sync()
}
```
