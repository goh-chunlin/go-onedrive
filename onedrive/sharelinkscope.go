// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package onedrive

// ShareLinkType the possible values for the scope property of SharingLink
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/resources/sharinglink?view=odsp-graph-online#scope-options
type ShareLinkScope int

const (
	Anonymous ShareLinkScope = iota
	Organization
)

func (shareLinkScope ShareLinkScope) toString() string {
	return [...]string{"anonymous", "organization"}[shareLinkScope]
}
