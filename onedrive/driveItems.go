// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package onedrive

import (
	"context"
	"errors"
	"net/url"
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

// NewFolderCreationRequest represents the information needed of a new OneDrive folder to be created.
type NewFolderCreationRequest struct {
	FolderName  string      `json:"name"`
	FolderFacet FolderFacet `json:"folder"`
	Restriction string      `json:"@microsoft.graph.conflictBehavior"`
}

// FolderFacet represents one of the facets to create a new folder.
type FolderFacet struct {
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
	apiURL := "me/drive/items/" + url.PathEscape(folderId) + "/children"
	if folderId == "" {
		apiURL = "me/drive/root/children"
	}

	req, err := s.client.NewRequest("GET", apiURL, nil)
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

// List the items of a special folder in the default drive of the authenticated user.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/drive_get_specialfolder?view=odsp-graph-online#get-children-of-a-special-folder
func (s *DriveItemsService) ListSpecial(ctx context.Context, folderName DriveSpecialFolder) (*OneDriveDriveItemsResponse, error) {
	apiURL := "me/drive/special/" + url.PathEscape(folderName.toString()) + "/children"

	req, err := s.client.NewRequest("GET", apiURL, nil)
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

	apiURL := "me/drive/items/" + url.PathEscape(itemId)

	req, err := s.client.NewRequest("GET", apiURL, nil)
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

// Get an item from special folder in the default drive of the authenticated user.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/drive_get_specialfolder?view=odsp-graph-online
func (s *DriveItemsService) GetSpecial(ctx context.Context, folderName DriveSpecialFolder) (*DriveItem, error) {
	if folderName.toString() == "" {
		return nil, errors.New("Please specify which special folder to use.")
	}

	apiURL := "me/drive/special/" + url.PathEscape(folderName.toString())

	req, err := s.client.NewRequest("GET", apiURL, nil)
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

// Create a new folder in a drive of the authenticated user.
// If there is already a folder in the same OneDrive directory with the same name,
// OneDrive will choose a new name for the folder while creating it.
//
// If driveId is empty, it means the selected drive will be the default drive of
// the authenticated user.
//
// If parentFolderName is empty, it means the new folder will be created at
// the root of the default drive.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/driveitem_post_children?view=odsp-graph-online
func (s *DriveItemsService) CreateNewFolder(ctx context.Context, driveId string, parentFolderName string, folderName string) (*DriveItem, error) {
	if folderName == "" {
		return nil, errors.New("Please provide the folder name.")
	}

	if parentFolderName == "" {
		parentFolderName = "root"
	}

	apiURL := "me/drive/items/" + url.PathEscape(parentFolderName) + "/children"
	if driveId != "" {
		apiURL = "me/drives/" + url.PathEscape(driveId) + "/items/" + url.PathEscape(parentFolderName) + "/children"
	}

	folderFacet := &FolderFacet{}

	newFolder := &NewFolderCreationRequest{
		FolderName:  folderName,
		FolderFacet: *folderFacet,
		Restriction: "rename",
	}

	req, err := s.client.NewRequest("POST", apiURL, newFolder)
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
