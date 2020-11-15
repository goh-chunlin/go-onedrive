// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package onedrive

import (
	"context"
	"errors"
)

// DriveItemsService handles communication with the drive items related methods of the OneDrive API.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/resources/driveitem?view=odsp-graph-online
type DriveItemsService service

// OneDriveDriveItemsResponse represents the JSON object returned by the OneDrive API.
type OneDriveDriveItemsResponse struct {
	ODataContext string       `json:"@odata.context"`
	Count        int          `json:"@odata.count"`
	DriveItems   []*DriveItem `json:"value"`
}

// DriveItem represents a OneDrive drive item.
type DriveItem struct {
	Name        string         `json:"name"`
	Id          string         `json:"id"`
	DownloadURL string         `json:"@microsoft.graph.downloadUrl"`
	Description string         `json:"description"`
	Audio       *OneDriveAudio `json:"audio"`
	Image       *OneDriveImage `json:"image"`
}

// OneDriveAudio represents the audio metadata of a OneDrive drive item which is an audio.
type OneDriveAudio struct {
	Title       string `json:"title"`
	Album       string `json:"album"`
	AlbumArtist string `json:"albumArtist"`
	Duration    int    `json:"duration"`
}

// OneDriveAudio represents the image metadata of a OneDrive drive item which is an image.
type OneDriveImage struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// List the items of a folder in the default drive of the authenticated user.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/resources/driveitem?view=odsp-graph-online
func (s *DriveItemsService) List(ctx context.Context, folderId string) (*OneDriveDriveItemsResponse, error) {
	apiURL := "me/drive/items/" + folderId + "/children"
	if folderId == "" {
		apiURL = "me/drive/root/children"
	}

	req, err := s.client.NewRequest("GET", apiURL)
	if err != nil {
		return nil, err
	}

	var oneDriveResponse *OneDriveDriveItemsResponse
	err = s.client.Do(ctx, req, &oneDriveResponse)
	if err != nil {
		return nil, err
	}

	return oneDriveResponse, nil
}

// Get an item in the default drive of the authenticated user.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/driveitem_get?view=odsp-graph-online
func (s *DriveItemsService) Get(ctx context.Context, itemId string) (*DriveItem, error) {
	if itemId == "" {
		return nil, errors.New("Please provide the Item ID of the item.")
	}

	apiURL := "me/drive/items/" + itemId

	req, err := s.client.NewRequest("GET", apiURL)
	if err != nil {
		return nil, err
	}

	var driveItem *DriveItem
	err = s.client.Do(ctx, req, &driveItem)
	if err != nil {
		return nil, err
	}

	return driveItem, nil
}
