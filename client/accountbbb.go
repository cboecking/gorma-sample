package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateAccountbbbPayload is the accountbbb create action payload.
type CreateAccountbbbPayload struct {
	// Name of accountbbb
	Name string `form:"name" json:"name" xml:"name"`
}

// CreateAccountbbbPath computes a request path to the create action of accountbbb.
func CreateAccountbbbPath() string {
	return fmt.Sprintf("/cellar-aaa/accountbbbs")
}

// Create new accountbbb
func (c *Client) CreateAccountbbb(ctx context.Context, path string, payload *CreateAccountbbbPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateAccountbbbRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateAccountbbbRequest create the request corresponding to the create action endpoint of the accountbbb resource.
func (c *Client) NewCreateAccountbbbRequest(ctx context.Context, path string, payload *CreateAccountbbbPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteAccountbbbPath computes a request path to the delete action of accountbbb.
func DeleteAccountbbbPath(accountbbbID int) string {
	return fmt.Sprintf("/cellar-aaa/accountbbbs/%v", accountbbbID)
}

// DeleteAccountbbb makes a request to the delete action endpoint of the accountbbb resource
func (c *Client) DeleteAccountbbb(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteAccountbbbRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteAccountbbbRequest create the request corresponding to the delete action endpoint of the accountbbb resource.
func (c *Client) NewDeleteAccountbbbRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListAccountbbbPath computes a request path to the list action of accountbbb.
func ListAccountbbbPath() string {
	return fmt.Sprintf("/cellar-aaa/accountbbbs")
}

// Retrieve all accountbbbs.
func (c *Client) ListAccountbbb(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListAccountbbbRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListAccountbbbRequest create the request corresponding to the list action endpoint of the accountbbb resource.
func (c *Client) NewListAccountbbbRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowAccountbbbPath computes a request path to the show action of accountbbb.
func ShowAccountbbbPath(accountbbbID int) string {
	return fmt.Sprintf("/cellar-aaa/accountbbbs/%v", accountbbbID)
}

// Retrieve accountbbb with given id.
func (c *Client) ShowAccountbbb(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowAccountbbbRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowAccountbbbRequest create the request corresponding to the show action endpoint of the accountbbb resource.
func (c *Client) NewShowAccountbbbRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateAccountbbbPayload is the accountbbb update action payload.
type UpdateAccountbbbPayload struct {
	// Name of accountbbb
	Name string `form:"name" json:"name" xml:"name"`
}

// UpdateAccountbbbPath computes a request path to the update action of accountbbb.
func UpdateAccountbbbPath(accountbbbID int) string {
	return fmt.Sprintf("/cellar-aaa/accountbbbs/%v", accountbbbID)
}

// Change accountbbb name
func (c *Client) UpdateAccountbbb(ctx context.Context, path string, payload *UpdateAccountbbbPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateAccountbbbRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateAccountbbbRequest create the request corresponding to the update action endpoint of the accountbbb resource.
func (c *Client) NewUpdateAccountbbbRequest(ctx context.Context, path string, payload *UpdateAccountbbbPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
