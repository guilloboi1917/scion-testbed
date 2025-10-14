package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Defines all handler function for ping commands

// StartPing issues a POST request to start a ping at src: from to dst: to with count: count
//
// Parameters:
// - dst: destination IP address (string)
// - count: number of ping packets to send (int) (-1 for no count)
//
// The scion node is expected to be running and reachable at c.baseURL
// Returns:
// - resp: pointer to http.Response containing the server's response
// - err: error if the request fails, otherwise nil
func (c *Client) StartPing(dstURL string, count int) (resp *http.Response, err error) {
	req := PingStartRequest{
		Dst: dstURL,
	}

	// Only set count if > 0
	if count > 0 {
		req.Count = &count
	}

	// Marshal request to JSON
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// Start new ping request
	resp, err = c.client.Post(c.baseURL+PingStartRoute, "application/json", bytes.NewBuffer(body))
	return
}

// StopPing issues a POST request to stop a ping at dst
//
// Parameters:
// - dst: destination IP address (string)
// - count: number of ping packets to send (int)
//
// Returns:
// - resp: pointer to http.Response containing the server's response
// - err: error if the request fails, otherwise nil
func (c *Client) StopPing() (resp *http.Response, err error) {
	// Start stop ping request
	resp, err = c.client.Post(c.baseURL+PingStopRoute, "application/json", nil)
	return
}

// ResultsPing issues a GET request to get the available ping results on a host
//
// Returns:
// - resp: pointer to http.Response containing the server's response
// - err: error if the request fails, otherwise nil
func (c *Client) GetResultsPing() (resp *http.Response, err error) {
	resp, err = c.client.Get(c.baseURL + PingListAvailableRoute)
	return
}

// TODO Add Doc
func (c *Client) StatusPing() (resp *http.Response, err error) {
	resp, err = c.client.Get(c.baseURL + PingStatusRoute)
	return
}
