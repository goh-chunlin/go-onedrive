// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package onedrive

import (
	"context"
)

type UserService service

// User represents an user in Microsoft Live.
type User struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	Email       string `json:"mail"`
	
}

// OneDrive API docs: https://learn.microsoft.com/en-us/graph/api/user-get?view=graph-rest-1.0&tabs=http
func (s *UserService) GetCurrentUserDetails(ctx context.Context) (*User, error) {
	apiURL := "me"

	req, err := s.client.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	var oneDriveResponse *User
	err = s.client.Do(ctx, req, false, &oneDriveResponse)
	if err != nil {
		return nil, err
	}

	return oneDriveResponse, nil
}
