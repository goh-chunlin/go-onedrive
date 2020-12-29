// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package onedrive

import (
	"bytes"
	"context"
	"errors"
	"net/url"
	"os"

	"github.com/h2non/filetype"
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
	FolderName  string `json:"name"`
	FolderFacet Facet  `json:"folder"`
	Restriction string `json:"@microsoft.graph.conflictBehavior"`
}

// Facet represents one of the facets for a folder or file.
type Facet struct {
}

// MoveItemRequest represents the information needed of moving an item in OneDrive.
type MoveItemRequest struct {
	ParentFolder ParentReference `json:"parentReference"`
}

// ParentReference represents the information of a folder in OneDrive.
type ParentReference struct {
	Id      string `json:"id"`
	Path    string `json:"path"`
	DriveId string `json:"driveId"`
}

// MoveItemResponse represents the JSON object returned by the OneDrive API after moving an item.
type MoveItemResponse struct {
	Id           string          `json:"id"`
	Name         string          `json:"name"`
	ParentFolder ParentReference `json:"parentReference"`
}

// RenameItemRequest represents the information needed of renaming an item in OneDrive.
type RenameItemRequest struct {
	Name string `json:"name"`
}

// RenameItemResponse represents the JSON object returned by the OneDrive API after renaming an item.
type RenameItemResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	File Facet  `json:"file"`
}

// CopyItemRequest represents the information needed of copying an item in OneDrive.
type CopyItemRequest struct {
	Name         string          `json:"name"`
	ParentFolder ParentReference `json:"parentReference"`
}

// CopyItemResponse represents the JSON object returned by the OneDrive API after copying an item.
type CopyItemResponse struct {
	Location string `json:"location"`
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
	err = s.client.Do(ctx, req, false, &oneDriveResponse)
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
	err = s.client.Do(ctx, req, false, &oneDriveResponse)
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
	err = s.client.Do(ctx, req, false, &driveItem)
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
	err = s.client.Do(ctx, req, false, &driveItem)
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

	folderFacet := &Facet{}

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
	err = s.client.Do(ctx, req, false, &driveItem)
	if err != nil {
		return nil, err
	}

	return driveItem, nil
}

// Delete a drive item in a drive of the authenticated user.
// The deleted item will be moved to the Recycle Bin instead of getting permanently deleted.
//
// If driveId is empty, it means the selected drive will be the default drive of
// the authenticated user.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/driveitem_delete?view=odsp-graph-online
func (s *DriveItemsService) Delete(ctx context.Context, driveId string, itemId string) (*DriveItem, error) {
	if itemId == "" {
		return nil, errors.New("Please provide the Item ID of the item to be deleted.")
	}

	apiURL := "me/drive/items/" + url.PathEscape(itemId)
	if driveId != "" {
		apiURL = "me/drives/" + url.PathEscape(driveId) + "/items/" + url.PathEscape(itemId)
	}

	req, err := s.client.NewRequest("DELETE", apiURL, nil)
	if err != nil {
		return nil, err
	}

	var driveItem *DriveItem
	err = s.client.Do(ctx, req, false, &driveItem)
	if err != nil {
		return nil, err
	}

	return driveItem, nil
}

// Move a drive item to a new parent folder in a drive of the authenticated user.
//
// When moving an item to the root of a drive, we cannot use "root" as the destinationParentFolderId.
// Instead, we need to provide the actual ID of the root.
//
// If driveId is empty, it means the selected drive will be the default drive of
// the authenticated user.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/driveitem_move?view=odsp-graph-online
func (s *DriveItemsService) Move(ctx context.Context, driveId string, itemId string, destinationParentFolderId string) (*MoveItemResponse, error) {
	if itemId == "" {
		return nil, errors.New("Please provide the Item ID of the item to be moved.")
	}

	if destinationParentFolderId == "" {
		return nil, errors.New("Please provide the destination, i.e. the ID of the new parent folder for the item.")
	}

	destinationParentFolder := &ParentReference{
		Id: destinationParentFolderId,
	}

	targetParentFolder := &MoveItemRequest{
		ParentFolder: *destinationParentFolder,
	}

	apiURL := "me/drive/items/" + url.PathEscape(itemId)
	if driveId != "" {
		apiURL = "me/drives/" + url.PathEscape(driveId) + "/items/" + url.PathEscape(itemId)
	}

	req, err := s.client.NewRequest("PATCH", apiURL, targetParentFolder)
	if err != nil {
		return nil, err
	}

	var response *MoveItemResponse
	err = s.client.Do(ctx, req, false, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Rename a drive item in a drive of the authenticated user.
//
// If driveId is empty, it means the selected drive will be the default drive of
// the authenticated user.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/driveitem_update?view=odsp-graph-online
func (s *DriveItemsService) Rename(ctx context.Context, driveId string, itemId string, newItemName string) (*RenameItemResponse, error) {
	if itemId == "" {
		return nil, errors.New("Please provide the Item ID of the item to be moved.")
	}

	if newItemName == "" {
		return nil, errors.New("Please provide a new name for the item.")
	}

	newNameRequest := &RenameItemRequest{
		Name: newItemName,
	}

	apiURL := "me/drive/items/" + url.PathEscape(itemId)
	if driveId != "" {
		apiURL = "me/drives/" + url.PathEscape(driveId) + "/items/" + url.PathEscape(itemId)
	}

	req, err := s.client.NewRequest("PATCH", apiURL, newNameRequest)
	if err != nil {
		return nil, err
	}

	var response *RenameItemResponse
	err = s.client.Do(ctx, req, false, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Copy a drive item to a new parent item or with a new name in a drive of the authenticated user.
//
// If sourceDriveId or destinationDriveId is empty, it means the selected drive will be the default drive of
// the authenticated user.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/driveitem_copy?view=odsp-graph-online
func (s *DriveItemsService) Copy(ctx context.Context, sourceDriveId string, itemId string,
	destinationDriveId string, destinationFolderId string, newItemName string) (*CopyItemResponse, error) {
	if itemId == "" {
		return nil, errors.New("Please provide the Item ID of the item to be copied.")
	}

	if destinationFolderId == "" {
		return nil, errors.New("Please provide the destination, i.e. the ID of the new parent folder for the item.")
	}

	if newItemName == "" {
		return nil, errors.New("Please provide the name of the new item after the copy is done. OneDrive will reject item name which already exists in destination.")
	}

	if destinationDriveId == "" {
		reqDefaultDriveInfo, err := s.client.NewRequest("GET", "me/drive", nil)
		if err != nil {
			return nil, err
		}

		var defaultDrive *Drive
		err = s.client.Do(ctx, reqDefaultDriveInfo, false, &defaultDrive)
		if err != nil {
			return nil, err
		}

		destinationDriveId = defaultDrive.Id
	}

	destinationParentFolder := &ParentReference{
		Id:      destinationFolderId,
		DriveId: destinationDriveId,
	}

	copyItemRequest := &CopyItemRequest{
		ParentFolder: *destinationParentFolder,
		Name:         newItemName,
	}

	apiURL := "me/drive/items/" + url.PathEscape(itemId) + "/copy"
	if sourceDriveId != "" {
		apiURL = "me/drives/" + url.PathEscape(sourceDriveId) + "/items/" + url.PathEscape(itemId) + "/copy"
	}

	req, err := s.client.NewRequest("POST", apiURL, copyItemRequest)
	if err != nil {
		return nil, err
	}

	var response *CopyItemResponse
	err = s.client.Do(ctx, req, false, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UploadNewFile is to upload a file to a drive of the authenticated user.
//
// If driveId is empty, it means the selected drive will be the default drive of
// the authenticated user.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/driveitem_put_content?view=odsp-graph-online#http-request-to-upload-a-new-file
func (s *DriveItemsService) UploadNewFile(ctx context.Context, driveId string, destinationParentFolderId string, localFilePath string) (*DriveItem, error) {
	if destinationParentFolderId == "" {
		return nil, errors.New("Please provide the destination, i.e. the ID of the parent folder for this new item.")
	}

	if localFilePath == "" {
		return nil, errors.New("Please provide the path to the file on local.")
	}

	file, err := os.Open(localFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if fileInfo.IsDir() {
		return nil, errors.New("Only file is allowed to be uploaded here.")
	}

	fileSize := fileInfo.Size()

	fileName := fileInfo.Name()

	apiURL := "me/drive/items/" + url.PathEscape(destinationParentFolderId) + ":/" + url.PathEscape(fileName) + ":/content"
	if driveId != "" {
		apiURL = "me/drives/" + url.PathEscape(driveId) + "/items/" + url.PathEscape(destinationParentFolderId) + ":/" + url.PathEscape(fileName) + ":/content"
	}

	buffer := make([]byte, fileSize)
	file.Read(buffer)
	fileReader := bytes.NewReader(buffer)

	fileType, _ := filetype.Match(buffer)

	req, err := s.client.NewFileUploadRequest(apiURL, fileType.MIME.Value, fileReader)
	if err != nil {
		return nil, err
	}

	var response *DriveItem
	err = s.client.Do(ctx, req, false, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
