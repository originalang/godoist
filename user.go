package togoist

import (
	"encoding/json"
	"fmt"
)

const BaseREST string = "https://beta.todoist.com/API/v8/"

type User struct {
	APIKey string
}

func NewUser(apiKey string) *User {
	return &User{apiKey}
}

func (u *User) Projects() []Project {
	// submit GET request to retrieve list of projects
	resp := request("GET", "projects", u.APIKey)
	defer resp.Body.Close()

	contents := readResponse(resp)

	pList := make([]Project, 0)
	json.Unmarshal(contents, &pList)

	return pList
}

func (u *User) AddProject(name string) Project {
	// prepare data string that will be used as the request body
	proj := fmt.Sprintf(`{"name":"%s"}`, name)

	// submit POST request to send data to todoist
	resp := request("POST", "projects", u.APIKey, proj)
	defer resp.Body.Close()

	contents := readResponse(resp)

	var p Project
	json.Unmarshal(contents, &p)

	return p
}

func (u *User) DeleteProject(project Project) {
	resp := request("DELETE", fmt.Sprintf("projects/%d", project.Id), u.APIKey)
	defer resp.Body.Close()
}

func (u *User) UpdateProject(project Project) {
	updatedProj := fmt.Sprintf(`{"name":"%s"}`, project.Name)
	resp := request("POST", fmt.Sprintf("projects/%d", project.Id), u.APIKey, updatedProj)
	defer resp.Body.Close()
}

func (u *User) Tasks() []Task {
	resp := request("GET", "tasks", u.APIKey)
	defer resp.Body.Close()

	contents := readResponse(resp)

	tList := make([]Task, 0)
	json.Unmarshal(contents, &tList)

	return tList
}

func (u *User) AddTask(content, due string, priority ...int) Task {

	// set priority based on input
	prty := 1
	if len(priority) > 0 {
		prty = priority[0]
	}

	task := fmt.Sprintf(`{"content":"%s", "due_string":"%s", "priority":%d}`, content, due, prty)

	resp := request("POST", "tasks", u.APIKey, task)
	defer resp.Body.Close()

	contents := readResponse(resp)

	var t Task
	json.Unmarshal(contents, &t)

	return t
}
