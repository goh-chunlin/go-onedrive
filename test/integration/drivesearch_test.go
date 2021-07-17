// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package integration

import (
	"fmt"
	"testing"
)

func TestDriveSearch_Search(t *testing.T) {
	ctx, client := setup()

	searchDriveItems, err := client.DriveSearch.SearchAll(ctx, "Shana")
	if err != nil {
		t.Errorf("Error: %v\n", err)
		return
	}
	for _, driveItem := range searchDriveItems.DriveItems {
		fmt.Printf("Results: %v\n", driveItem.Name)
	}
}

func TestDriveSearch_SearchWithApostrophe(t *testing.T) {
	ctx, client := setup()

	searchDriveItems, err := client.DriveSearch.Search(ctx, "Rabbit's")
	if err != nil {
		t.Errorf("Error: %v\n", err)
		return
	}
	for _, driveItem := range searchDriveItems.DriveItems {
		fmt.Printf("Results: %v\n", driveItem.Name)
	}
}
