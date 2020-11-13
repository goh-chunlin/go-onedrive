// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package onedrive

// User represents an user in Microsoft Live.
type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}
