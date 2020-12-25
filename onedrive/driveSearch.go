// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package onedrive

import (
	"context"
	"fmt"
	"strings"
)

// DriveSearchService handles communication with the drive items searching related methods of the OneDrive API.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/driveitem_search?view=odsp-graph-online
type DriveSearchService service

// OneDriveDriveSearchResponse represents the JSON object returned by the OneDrive API.
type OneDriveDriveSearchResponse struct {
	ODataContext string       `json:"@odata.context"`
	DriveItems   []*DriveItem `json:"value"`
}

// Search the items in the default drive of the authenticated user.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/driveitem_search?view=odsp-graph-online#request
func (s *DriveSearchService) Search(ctx context.Context, query string) (*OneDriveDriveSearchResponse, error) {
	// For requests that use single quotes, if any parameter values
	// also contain single quotes, those must be double escaped; otherwise,
	// the request will fail due to invalid syntax.
	//
	// Reference: https://docs.microsoft.com/en-us/graph/query-parameters
	query = strings.Replace(query, "'", "''", -1)

	apiURL := fmt.Sprintf("me/drive/root/search(q='%v')", query)

	req, err := s.client.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	var oneDriveResponse *OneDriveDriveSearchResponse
	err = s.client.Do(ctx, req, &oneDriveResponse)
	if err != nil {
		return nil, err
	}

	return oneDriveResponse, nil
}

// Search the items in the default drive of the authenticated user as well as items shared with the user.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/driveitem_search?view=odsp-graph-online#searching-for-items-a-user-can-access
func (s *DriveSearchService) SearchAll(ctx context.Context, query string) (*OneDriveDriveSearchResponse, error) {
	query = strings.Replace(query, "'", "''", -1)

	apiURL := fmt.Sprintf("me/drive/search(q='%v')", query)

	req, err := s.client.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	var oneDriveResponse *OneDriveDriveSearchResponse
	err = s.client.Do(ctx, req, &oneDriveResponse)
	if err != nil {
		return nil, err
	}

	return oneDriveResponse, nil
}
