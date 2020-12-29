package integration

import (
	"fmt"
	"testing"
)

func TestDriveSearch_Search(t *testing.T) {
	ctx, client := setup()

	searchDriveItems, err := client.DriveSearch.SearchAll(ctx, "Shana")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
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
