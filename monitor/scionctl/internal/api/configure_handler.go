package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Defines all handler function for configure commands

// ConfigureASList issues a POST request to configure AS blacklist on a node
//
// Parameters:
// - asList: list of AS identifiers in format "ffaa:1:<number>" ([]string)
//
// Returns:
// - resp: pointer to http.Response containing the server's response
// - err: error if the request fails, otherwise nil
func (c *Client) ConfigureASList(asList []string) (resp *http.Response, err error) {
	req := ConfigureASListRequest{
		ASList: asList,
	}

	// Marshal request to JSON
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// Send POST request
	resp, err = c.client.Post(c.baseURL+ConfigASListRoute, "application/json", bytes.NewBuffer(body))
	return
}
