package togoist

import (
	"bytes"
	"io"
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
func request(method, urlAppend, apiKey string, data ...string) *http.Response {

	client := &http.Client{}

	// parse and return the (optional) data
	// returns nil if no data provided (for GET requests, etc.)
	body := parseData(data)

	req, e := http.NewRequest(method, BaseREST+urlAppend, body)
	checkErr(e)

	// this header is shared by all request methods
	req.Header.Add("Authorization", "Bearer "+apiKey)

	// additional headers needed for different request methods
	switch method {
	case "POST":
		req.Header.Add("Content-Type", "application/json")
	}

	// submit request
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

func parseData(data []string) io.Reader {
	var httpBody io.Reader

	if len(data) > 0 {
		content := []byte(data[0])
		httpBody = bytes.NewBuffer(content)
	}

	return httpBody
}
