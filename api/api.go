package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// APIResponse is the Struct representation of a Json Response
type APIResponse struct {
	Header APIHeader       `json:"header"`
	Body   json.RawMessage `json:"body"`
}

// APIHeader is the Struct representation of the Header of a APIResponse
type APIHeader struct {
	ID         string `json:"id"`
	Status     string `json:"status"`
	Servertime int    `json:"servertime"`
	Action     string `json:"action"`
	Message    string `json:"message"`
	URL        string `json:"url"`
	Code       int    `json:"code"`
}

// DoCustomRequest Executes a Custom Request and returns a APIResponse
func (c *Client) DoCustomRequest(ctx context.Context, method, path, version string, body interface{}, opts interface{}) (*APIResponse, error) {
	_, response, err := c.DoCustomRequestAndReturnRawResponse(ctx, method, path, version, body, opts)
	return response, err
}

// DoCustomRequestAndReturnRawResponse Executes a Custom Request and returns a APIResponse and the Raw HTTP Response
func (c *Client) DoCustomRequestAndReturnRawResponse(ctx context.Context, method, path, version string, body interface{}, opts interface{}) (*http.Response, *APIResponse, error) {
	u, err := addOptions(path, version, opts)
	if err != nil {
		return nil, nil, fmt.Errorf("Adding Request Options: %w", err)
	}

	req, err := c.newRequest(method, u, body)
	if err != nil {
		return nil, nil, fmt.Errorf("Creating New Request: %w", err)
	}

	var res APIResponse
	r, err := c.do(ctx, req, &res)
	if err != nil {
		return r, &res, fmt.Errorf("Doing Request: %w", err)
	}

	if res.Header.Status == "success" {
		return r, &res, nil
	} else if res.Header.Status == "error" {
		return r, &res, fmt.Errorf("%w: Message: %v, Body: %v", ErrAPIResponseErrorStatusCode, res.Header.Message, string(res.Body))
	} else {
		return r, &res, fmt.Errorf("%w: Message: %v, Body: %v", ErrAPIResponseUnknownStatusCode, res.Header.Message, string(res.Body))
	}
}
