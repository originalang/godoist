package togoist

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func decodeResponse(resp *http.Response) Response {
	content, _ := ioutil.ReadAll(resp.Body)

	var decoded Response
	json.Unmarshal(content, &decoded)

	return decoded
}

func projectToMap(p *Project) map[string]interface{} {
	proj, _ := json.Marshal(p)

	var newMap map[string]interface{}
	json.Unmarshal(proj, &newMap)

	// the unmarshaling process does not parse the id field correctly
	// this is done to preserve the correct id for the project
	newMap["id"] = p.Id

	return newMap
}
