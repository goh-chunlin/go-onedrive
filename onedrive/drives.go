// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package onedrive

import (
	"context"
	"net/url"
)

// DrivesService handles communication with the drives related methods of the OneDrive API.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/resources/drive?view=odsp-graph-online
type DrivesService service

// OneDriveDrivesResponse represents the JSON object containing drive list returned by the OneDrive API.
type OneDriveDrivesResponse struct {
	ODataContext string   `json:"@odata.context"`
	Drives       []*Drive `json:"value"`
}

// Drive represents a OneDrive drive.
type Drive struct {
	Id        string      `json:"id"`
	DriveType string      `json:"driveType"`
	Owner     *Owner      `json:"owner"`
	Quota     *DriveQuota `json:"quota"`
}

// DriveQuota represents the usage quota of a drive.
type DriveQuota struct {
	Used      int    `json:"used"`
	Deleted   int    `json:"deleted"`
	Remaining int    `json:"remaining"`
	Total     int    `json:"total"`
	State     string `json:"state"`
}

// Get a specified drive (or default drive if driveId is empty) of the authenticated user.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/drive_get?view=odsp-graph-online
func (s *DrivesService) Get(ctx context.Context, driveId string) (*Drive, error) {
	apiURL := "me/drives/" + url.PathEscape(driveId)
	if driveId == "" {
		apiURL = "me/drive"
	}

	req, err := s.client.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	var defaultDrive *Drive
	err = s.client.Do(ctx, req, false, &defaultDrive)
	if err != nil {
		return nil, err
	}

	return defaultDrive, nil
}

// List the drives of the authenticated user.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/drive_list?view=odsp-graph-online
func (s *DrivesService) List(ctx context.Context) (*OneDriveDrivesResponse, error) {
	req, err := s.client.NewRequest("GET", "me/drives", nil)
	if err != nil {
		return nil, err
	}

	var oneDriveResponse *OneDriveDrivesResponse
	err = s.client.Do(ctx, req, false, &oneDriveResponse)
	if err != nil {
		return nil, err
	}

	return oneDriveResponse, nil
}
