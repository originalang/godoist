package togoist

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const BaseREST string = "https://beta.todoist.com/API/v8/"

type User struct {
	APIKey string
}

func NewUser(apiKey string) *User {
	return &User{apiKey}
}

func (u *User) Projects() []Project {

	client := &http.Client{}

	req, _ := http.NewRequest("GET", BaseREST+"projects", nil)
	req.Header.Set("Authorization", "Bearer "+u.APIKey)

	resp, e := client.Do(req)
	err(e)

	defer resp.Body.Close()

	contents, e := ioutil.ReadAll(resp.Body)
	err(e)

	p := make([]Project, 0)
	json.Unmarshal(contents, &p)

	return p
}

func err(e error) {
	if e != nil {
		panic(e)
	}
}
