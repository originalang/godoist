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