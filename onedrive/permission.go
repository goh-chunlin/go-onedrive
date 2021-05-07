package onedrive

import (
	"context"
	"fmt"
	"net/http"
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
func (s *PermissionService) CreateShareLink(ctx context.Context, itemID, permissionType, permissionScope string) (*Permission, error) {
	apiURL := fmt.Sprintf("me/drive/items/%s/createLink", itemID)

	body := &CreateShareLinkRequest{Type: permissionType, Scope: permissionScope}
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
// OneDrive API docs:  https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/driveitem_list_permissions?view=odsp-graph-online
func (s *PermissionService) List(ctx context.Context, itemID string) ([]Permission, error) {
	apiURL := fmt.Sprintf("me/drive/items/%s/permissions", itemID)

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
