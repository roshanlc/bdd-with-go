package main

import (
	"net/http"
)

// ApiResponse struct wraps http.Response
// and error
type ApiResponse struct {
	*http.Response       // embedding http.Response
	err            error //error
}

// getReq is a helper function that wraps
// multiple steps to send a GET request and
// returns it's reponse
func getReq(url string) *ApiResponse {

	var resp ApiResponse

	// method
	method := "GET"
	// new http client
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		resp.err = err
		return &resp
	}
	// Set Header
	req.Header.Add("Content-Type", "application/json")

	// Send the request
	res, err := client.Do(req)

	// Check for error
	if err != nil {
		resp.err = err
		return &resp
	}
	// set the reponse and request
	resp.Response = res

	return &resp
}
