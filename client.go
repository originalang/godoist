package togoist

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
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
	Projects      map[string]Project
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

	// map projects to a key-value pair
	projectMap := make(map[string]Project)
	for _, proj := range resp.Projects {
		projectMap[proj.Name] = proj
	}
	c.Projects = projectMap

	c.Items = resp.Items
	sort.SliceStable(c.Items, func(i, j int) bool { return c.Items[i].ItemOrder < c.Items[j].ItemOrder })
}

func (c *Client) request() *http.Response {

	body := c.encodeBody()

	req, err := http.NewRequest("POST", endpoint, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	checkErr(err)

	resp, err := c.HTTPClient.Do(req)
	checkErr(err)
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

func (c *Client) DeleteProjects(ids []int64) {
	cmd := NewCommand("project_delete", map[string]interface{}{"ids": ids})
	c.setAttributes(`"projects"`, []string{cmd.Stringify()})

	c.performRequest()
}

func (c *Client) ArchiveProjects(ids []int64) {
	cmd := NewCommand("project_archive", map[string]interface{}{"ids": ids})
	c.setAttributes(`"projects"`, []string{cmd.Stringify()})

	c.performRequest()
}

func (c *Client) UnarchiveProjects(ids []int64) {
	cmd := NewCommand("project_unarchive", map[string]interface{}{"ids": ids})
	c.setAttributes(`"projects"`, []string{cmd.Stringify()})

	c.performRequest()
}

func (c *Client) AddItem(projectId int64, content string, indent int, dueDate string) Item {
	cmd := NewCommand("item_add", map[string]interface{}{"project_id": projectId, "content": content, "indent": indent, "date_string": dueDate})
	c.setAttributes(`"items"`, []string{cmd.Stringify()})

	resp := c.performRequest()

	return resp.Items[0]
}

func (c *Client) UpdateItem(item Item) Item {
	fmt.Printf("%+v", itemToMap(&item))
	cmd := NewCommand("item_update", itemToMap(&item))
	c.setAttributes(`"items"`, []string{cmd.Stringify()})

	resp := c.performRequest()

	return resp.Items[0]
}

func (c *Client) DeleteItems(ids []int64) {
	cmd := NewCommand("item_delete", map[string]interface{}{"ids": ids})
	c.setAttributes(`"items"`, []string{cmd.Stringify()})

	c.performRequest()
}

func (c *Client) CompleteItems(ids []int64, toHistory bool) {
	var cmd *Command
	if toHistory {
		cmd = NewCommand("item_complete", map[string]interface{}{"ids": ids, "force_history": 1})
	} else {
		cmd = NewCommand("item_complete", map[string]interface{}{"ids": ids})
	}
	c.setAttributes(`"items"`, []string{cmd.Stringify()})

	c.performRequest()
}

func (c *Client) UncompleteItems(ids []int64) {
	cmd := NewCommand("item_complete", map[string]interface{}{"ids": ids})
	c.setAttributes(`"items"`, []string{cmd.Stringify()})

	c.performRequest()
}

// non-modifying functions

// retrieve a project's id by name
func GetProjectId(c *Client, name string) (int64, error) {
	proj, ok := c.Projects[name]

	if ok {
		return proj.Id, nil
	} else {
		return 0, errors.New("Project does not exist")
	}
}

func GetChildrenIds(c *Client, parentId int64) []int64 {
	var children []int64

	for _, itm := range c.Items {
		if itm.ParentId == parentId {
			children = append(children, itm.Id)
		}
	}

	return children
}
