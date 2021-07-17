package onedrive

import (
	"context"
	"errors"
	"net/http"
	"net/url"
)

// PermissionService handles permission settings of a drive item
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/resources/permission?view=odsp-graph-online
type PermissionService service

// Permission is the permission of a drive item.
type Permission struct {
	ID        string      `json:"id"`
	GrantedTo interface{} `json:"grantedTo"`
	Link      SharingLink `json:"link"`
	Roles     []string    `json:"roles"`
}

// CreateShareLinkRequest is the request for creating a share link.
type CreateShareLinkRequest struct {
	Type  string `json:"type"`  // The type of sharing link to create. Either view, edit, or embed.
	Scope string `json:"scope"` // Optional. The scope of link to create. Either anonymous or organization.
}

// SharingLink resource groups link-related data items into a single structure.
// Ref: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/resources/sharinglink?view=odsp-graph-online
type SharingLink struct {
	Type  string `json:"type"`  // The type of sharing link to create. Either view, edit, or embed.
	Scope string `json:"scope"` // Optional. The scope of link to create. Either anonymous or organization.
	URL   string `json:"webUrl"`
}

// CreateShareLink will create a new sharing link if the specified link type doesn't already exist for the calling application.
// If a sharing link of the specified type already exists for the app, the existing sharing link will be returned.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/driveitem_createlink?view=odsp-graph-online
func (s *PermissionService) CreateShareLink(ctx context.Context, itemId string, permissionType ShareLinkType, permissionScope ShareLinkScope) (*Permission, error) {
	apiURL := "me/drive/items/" + url.PathEscape(itemId) + "/createLink"

	body := &CreateShareLinkRequest{Type: permissionType.toString(), Scope: permissionScope.toString()}
	req, err := s.client.NewRequest(http.MethodPost, apiURL, body)
	if err != nil {
		return nil, err
	}

	var oneDriveResponse *Permission
	err = s.client.Do(ctx, req, false, &oneDriveResponse)
	if err != nil {
		return nil, err
	}

	return oneDriveResponse, nil
}

// ListPermissionsResponse is the response of list permissions of a drive item
type ListPermissionsResponse struct {
	Value []Permission `json:"value"`
}

// List lists the effective sharing permissions of on a DriveItem.
//
// OneDrive API docs:  https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/driveitem_list_permissions?view=odsp-graph-online
func (s *PermissionService) List(ctx context.Context, itemId string) ([]Permission, error) {
	apiURL := "me/drive/items/" + url.PathEscape(itemId) + "/permissions"

	req, err := s.client.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, err
	}

	var oneDriveResponse *ListPermissionsResponse
	err = s.client.Do(ctx, req, false, &oneDriveResponse)
	if err != nil {
		return nil, err
	}

	return oneDriveResponse.Value, nil
}

// Delete will delete a sharing permission from a file or folder.
// Only sharing permissions that are not inherited can be deleted. The inheritedFrom property must be null.
//
// If driveId is empty, it means the selected drive will be the default drive of
// the authenticated user.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/driveitem_delete?view=odsp-graph-online
func (s *PermissionService) Delete(ctx context.Context, driveId string, itemId string, permissionId string) error {
	if itemId == "" {
		return errors.New("Please provide the Item ID of the item to be deleted.")
	}

	apiURL := "me/drive/items/" + url.PathEscape(itemId)
	if driveId != "" {
		apiURL = "me/drives/" + url.PathEscape(driveId) + "/items/" + url.PathEscape(itemId)
	}

	apiURL += "/permissions/" + url.PathEscape(permissionId)

	req, err := s.client.NewRequest("DELETE", apiURL, nil)
	if err != nil {
		return err
	}

	var driveItem *DriveItem
	err = s.client.Do(ctx, req, false, &driveItem)
	if err != nil {
		return err
	}

	return nil
}
