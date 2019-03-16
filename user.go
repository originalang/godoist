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
