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
}

func NewClient(token string) *Client {
	return &Client{
		HTTPClient:    &http.Client{},
		Token:         token,
		SyncToken:     "*",
		ResourceTypes: `"projects"`,
	}
}

func (c *Client) NewRequest(method, urlPath string) *http.Response {

	body := c.encodeBody()

	req, _ := http.NewRequest(method, endpoint, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := c.HTTPClient.Do(req)
	return resp
}

func (c *Client) encodeBody() *strings.Reader {
	return strings.NewReader(fmt.Sprintf(`token=%s&sync_token=%s&resource_types=[%s]`, c.Token, c.SyncToken, c.ResourceTypes))
}

func (c *Client) setResourceTypes(resources string) {
	c.ResourceTypes = resources
}
