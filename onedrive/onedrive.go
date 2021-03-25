// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package onedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL  = "https://graph.microsoft.com/v1.0/"
	oneDriveBaseUrl = "https://api.onedrive.com/v1.0/"
)

type service struct {
	client *Client
}

// A Client manages communication with the OneDrive API.
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public OneDrive API. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the OneDrive API.
	Drives           *DrivesService
	DriveItems       *DriveItemsService
	DriveSearch      *DriveSearchService
	DriveAsyncJob    *DriveAsyncJobService
	DrivePermissions *PermissionService
}

// NewClient returns a new OneDrive API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL}

	c.common.client = c

	c.Drives = (*DrivesService)(&c.common)
	c.DriveItems = (*DriveItemsService)(&c.common)
	c.DriveSearch = (*DriveSearchService)(&c.common)
	c.DriveAsyncJob = (*DriveAsyncJobService)(&c.common)
	c.DrivePermissions = (*PermissionService)(&c.common)

	return c
}

// NewRequest creates an API request. A relative URL can be provided in relativeURL,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified WITHOUT a preceding slash.
func (c *Client) NewRequest(method, relativeURL string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not.", c.BaseURL)
	}

	apiUrl, err := c.BaseURL.Parse(relativeURL)
	if err != nil {
		return nil, err
	}

	if body != nil {
		jsonBody, err := json.Marshal(body)

		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest(method, apiUrl.String(), bytes.NewBuffer([]byte(jsonBody)))
		req.Header.Set("Content-Type", "application/json")

		return req, nil
	}

	// Create a new request using http
	req, err := http.NewRequest(method, apiUrl.String(), nil)

	return req, err
}

// NewFileUploadRequest creates an API request to upload files. A relative URL can be provided in relativeURL,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified WITHOUT a preceding slash.
func (c *Client) NewFileUploadRequest(relativeURL string, contentType string, fileReader *bytes.Reader) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not.", c.BaseURL)
	}

	if fileReader == nil {
		return nil, errors.New("Please provide the file reader.")
	}

	apiUrl, err := c.BaseURL.Parse(relativeURL)
	if err != nil {
		return nil, err
	}

	// Create a new request using http
	req, err := http.NewRequest("PUT", apiUrl.String(), fileReader)
	req.Header.Set("Content-Type", contentType)

	return req, err
}

// NewRequest creates an API request to OneDrive API directly with an absolute URL.
func (c *Client) NewRequestToOneDrive(method, absoluteUrl string, body interface{}) (*http.Request, error) {
	if !strings.HasPrefix(absoluteUrl, oneDriveBaseUrl) && !strings.HasPrefix(absoluteUrl, "/test-onedrive-api") {
		return nil, fmt.Errorf("The given URL %q is not a OneDrive API URL.", c.BaseURL)
	}

	if strings.HasPrefix(absoluteUrl, "/test-onedrive-api") {
		apiUrl, err := c.BaseURL.Parse(absoluteUrl)
		if err != nil {
			return nil, err
		}

		absoluteUrl = apiUrl.String()
	}

	if body != nil {
		jsonBody, err := json.Marshal(body)

		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest(method, absoluteUrl, bytes.NewBuffer([]byte(jsonBody)))
		req.Header.Set("Content-Type", "application/json")

		return req, nil
	}

	// Create a new request using http
	req, err := http.NewRequest(method, absoluteUrl, nil)

	return req, err
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by target, or returned as an
// error if an API error has occurred.
func (c *Client) Do(ctx context.Context, req *http.Request, isUsingPlainHttpClient bool, target interface{}) error {
	if ctx == nil {
		return errors.New("context must be non-nil")
	}
	req = req.WithContext(ctx)

	var err error
	var resp *http.Response

	if isUsingPlainHttpClient {
		httpClient := &http.Client{}
		resp, err = httpClient.Do(req)
	} else {
		resp, err = c.client.Do(req)
	}

	if err != nil {
		// If we got an error, and the context has been canceled, the error from the context is probably more useful.
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = sanitizeURL(url).String()
				return e
			}
		}

		return err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	locationHeader, isLocationHeaderExist := resp.Header["Location"]

	if resp.StatusCode == 202 && isLocationHeaderExist && len(responseBody) == 0 {

		var jsonStream = "{\"Location\": \"" + locationHeader[0] + "\"}"

		err = json.NewDecoder(strings.NewReader(jsonStream)).Decode(target)

	} else if resp.StatusCode != 204 {

		responseBodyReader := bytes.NewReader(responseBody)

		var oneDriveError *ErrorResponse
		json.NewDecoder(responseBodyReader).Decode(&oneDriveError)

		if oneDriveError.Error != nil {
			if oneDriveError.Error.InnerError != nil {
				return errors.New(oneDriveError.Error.Code + " - " + oneDriveError.Error.Message + " (" + oneDriveError.Error.InnerError.Date + ")")
			}

			return errors.New(oneDriveError.Error.Code + " - " + oneDriveError.Error.Message)
		}

		responseBodyReader = bytes.NewReader(responseBody)
		err = json.NewDecoder(responseBodyReader).Decode(target)

	}

	return err
}

// sanitizeURL redacts the client_secret parameter from the URL which may be exposed to the user.
func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("client_secret")) > 0 {
		params.Set("client_secret", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}
