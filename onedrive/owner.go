// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package onedrive

// Owner represents the owner of a OneDrive drive.
type Owner struct {
	User User `json:"user"`
}
