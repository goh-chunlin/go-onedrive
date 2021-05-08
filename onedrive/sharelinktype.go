// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package onedrive

// ShareLinkType the possible values for the type property of SharingLink
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/resources/sharinglink?view=odsp-graph-online#type-options
type ShareLinkType int

const (
	View ShareLinkType = iota
	Edit
	Embed
)

func (shareLinkType ShareLinkType) toString() string {
	return [...]string{"view", "edit", "embed"}[shareLinkType]
}
