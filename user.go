package togoist

import (
	"encoding/json"
)

const BaseREST string = "https://beta.todoist.com/API/v8/"

type User struct {
	APIKey string
}

func NewUser(apiKey string) *User {
	return &User{apiKey}
}

func (u *User) Projects() []Project {
	resp := request("GET", "projects", u.APIKey)
	defer resp.Body.Close()

	contents := readResponse(resp)

	pList := make([]Project, 0)
	json.Unmarshal(contents, &pList)

	return pList
}
