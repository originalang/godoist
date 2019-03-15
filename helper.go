package togoist

import (
	"io/ioutil"
	"net/http"
)

// Evaluates errors returned from other functions
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

// perform the specified request and return the response
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

// read the contents of an http response
func readResponse(resp *http.Response) []byte {
	contents, e := ioutil.ReadAll(resp.Body)
	checkErr(e)

	return contents
}
