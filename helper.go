package togoist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// test returned errors, and print 
// the error out if it exists
func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

// decode an http response by unmarshalling it
// into the Response struct
func decodeResponse(resp *http.Response) Response {
	content, _ := ioutil.ReadAll(resp.Body)

	var decoded Response
	json.Unmarshal(content, &decoded)

	return decoded
}

// convert a project struct to a map
func projectToMap(p *Project) map[string]interface{} {
	proj, _ := json.Marshal(p)

	var newMap map[string]interface{}
	json.Unmarshal(proj, &newMap)

	// the unmarshaling process does not parse the id field correctly
	// this is done to preserve the correct id for the project
	newMap["id"] = p.Id

	return newMap
}

// convert an item struct to a map
func itemToMap(item *Item) map[string]interface{} {
	proj, _ := json.Marshal(item)

	var newMap map[string]interface{}
	json.Unmarshal(proj, &newMap)

	// the unmarshaling process does not parse the id field correctly
	// this is done to preserve the correct id for the project
	newMap["id"] = item.Id
	newMap["project_id"] = item.ProjectId
	newMap["parent_id"] = item.ParentId
	newMap["responsible_uid"] = item.ResponsibleUserId
	newMap["user_id"] = item.UserId
	newMap["assigned_by_uid"] = item.AssignedBy

	// remove any keys that have blank values
	// doing this ensures that the values are passed correctly to the API
	for k, v := range newMap {
		if v == "" {
			delete(newMap, k)
		}
	}
	return newMap
}
