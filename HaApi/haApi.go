package haApi

import (
	"bytes"
	"log"
	"net/http"
)

const haUrl = "http://homeassistant.local:8123/api/"
const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiIwOGI4MzNlYTM1NmE0MmFlOWFkMDk1NTcyYTdiNDYwOCIsImlhdCI6MTczMTE2NDAzMCwiZXhwIjoyMDQ2NTI0MDMwfQ.PzrIUWGJYR3BKciZShxDhqKjjUn6xPvkdqEUupjmFv8"

func httpGet(url string) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	handleErr(err)

	req.Header.Add("Authorization", "Bearer "+token)
	return client.Do(req)
}

func httpPost(url string, payload []byte) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	handleErr(err)

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	return client.Do(req)
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
