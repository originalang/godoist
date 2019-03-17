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
	return strings.NewReader(fmt.Sprintf(`token=%s&sync_token=%s&resource_types=[%s]`, c.Token, c.SyncToken, c.ResourceTypes))
}

func (c *Client) decodeResponse(resp *http.Response) map[string]interface{} {
	content, _ := ioutil.ReadAll(resp.Body)

	var decoded map[string]interface{}
	json.Unmarshal(content, &decoded)

	c.SyncToken = decoded["sync_token"].(string)

	return decoded
}

func (c *Client) setResourceTypes(resources string) {
	c.ResourceTypes = resources
}

func (c *Client) Projects() interface{} {
	r := c.request()

	defer r.Body.Close()

	a := c.decodeResponse(r)

	return a["projects"]
}
