package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func setDefaultHeader(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
}

// Post is the standard http POST verb implementation
func Post(url string, data []byte) ([]byte, error) {

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	setDefaultHeader(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)

	return body, err
}
