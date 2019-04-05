package togoist

import (
	"fmt"
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
	User          User
	Projects      []Project
	Items         []Item
}

func NewClient(token string) *Client {
	return &Client{
		HTTPClient:    &http.Client{},
		Token:         token,
		SyncToken:     "*",
		ResourceTypes: `"projects"`,
	}
}

func (c *Client) Sync() {
	c.ResourceTypes = `"all"`
	c.SyncToken = "*"

	r := c.request()
	defer r.Body.Close()

	resp := decodeResponse(r)

	// update the sync token on the client
	c.SyncToken = resp.SyncToken

	// Perform sync
	c.User = resp.User
	c.Projects = resp.Projects
	c.Items = resp.Items
}

func (c *Client) request() *http.Response {

	body := c.encodeBody()

	req, _ := http.NewRequest("POST", endpoint, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := c.HTTPClient.Do(req)
	return resp
}

func (c *Client) performRequest() Response {
	r := c.request()
	defer r.Body.Close()

	d := decodeResponse(r)

	// update the sync token on the client
	c.SyncToken = d.SyncToken

	return d
}

func (c *Client) encodeBody() *strings.Reader {
	return strings.NewReader(fmt.Sprintf(`token=%s&sync_token=%s&resource_types=[%s]&commands=[%s]`, c.Token, c.SyncToken, c.ResourceTypes, strings.Join(c.Commands, ", ")))
}

func (c *Client) setAttributes(resources string, commands []string) {
	c.ResourceTypes = resources
	c.Commands = commands
}

func (c *Client) AddProject(name string, indent int) Project {
	cmd := NewCommand("project_add", map[string]interface{}{"name": name, "indent": indent})
	c.setAttributes(`"projects"`, []string{cmd.Stringify()})

	resp := c.performRequest()

	return resp.Projects[0]
}

func (c *Client) UpdateProject(p Project) Project {
	cmd := NewCommand("project_update", projectToMap(&p))
	c.setAttributes(`"projects"`, []string{cmd.Stringify()})

	resp := c.performRequest()

	return resp.Projects[0]
}
