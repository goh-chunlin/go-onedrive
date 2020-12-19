// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package onedrive

// DriveSpecialFolder indicates the pre-defined special folder in OneDrive
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/drive_get_specialfolder?view=odsp-graph-online#special-folder-names
type DriveSpecialFolder int

const (
	Documents DriveSpecialFolder = iota
	Photos
	CameraRoll
	AppRoot
	Music
)

func (specialFolder DriveSpecialFolder) toString() string {
	return [...]string{"documents", "photos", "cameraroll", "approot", "music"}[specialFolder]
}
