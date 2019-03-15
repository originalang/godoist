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

func request(method, urlAppend, apiKey string) *http.Response {

	client := &http.Client{}

	req, e := http.NewRequest(method, BaseREST+urlAppend, nil)
	checkErr(e)

	switch method {
	case "GET":
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}

	resp, e := client.Do(req)
	checkErr(e)

	return resp
}

func (u *User) Projects() []Project {
	resp := request("GET", "projects", u.APIKey)
	defer resp.Body.Close()

	contents, e := ioutil.ReadAll(resp.Body)
	checkErr(e)

	pList := make([]Project, 0)
	json.Unmarshal(contents, &pList)

	return pList
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
