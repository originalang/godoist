package togoist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const endpoint string = "https://todoist.com/api/v7/sync"

type Client struct {
	HTTPClient    *http.Client
	Token         string
	SyncToken     string
	ResourceTypes string
	Commands      []string
}

func NewClient(token string) *Client {
	return &Client{
		HTTPClient:    &http.Client{},
		Token:         token,
		SyncToken:     "*",
		ResourceTypes: `"projects"`,
	}
}

func (c *Client) request() *http.Response {

	body := c.encodeBody()

	req, _ := http.NewRequest("POST", endpoint, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := c.HTTPClient.Do(req)
	return resp
}

func (c *Client) encodeBody() *strings.Reader {
	return strings.NewReader(fmt.Sprintf(`token=%s&sync_token=%s&resource_types=[%s]&commands=[%s]`, c.Token, c.SyncToken, c.ResourceTypes, strings.Join(c.Commands, ", ")))
}

func (c *Client) decodeResponse(resp *http.Response) map[string]interface{} {
	content, _ := ioutil.ReadAll(resp.Body)

	var decoded map[string]interface{}
	json.Unmarshal(content, &decoded)

	c.SyncToken = decoded["sync_token"].(string)

	return decoded
}

func (c *Client) setAttributes(resources string, commands []string) {
	c.ResourceTypes = resources
	c.Commands = commands
}

func (c *Client) Projects() interface{} {
	r := c.request()

	defer r.Body.Close()

	d := c.decodeResponse(r)

	return d["projects"]
}

func (c *Client) AddProject(name string, indent int) interface{} {
	cmd := NewCommand("project_add", map[string]interface{}{"name": name, "indent": indent})
	c.setAttributes(`"projects"`, []string{cmd.Stringify()})
	r := c.request()

	defer r.Body.Close()

	d := c.decodeResponse(r)

	return d["projects"]

}
