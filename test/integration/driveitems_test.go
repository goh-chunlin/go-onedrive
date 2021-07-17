// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package integration

import (
	"fmt"
	"testing"

	"github.com/goh-chunlin/go-onedrive/onedrive"
)

func TestDriveItems_GetItemsInDefaultDriveRoot(t *testing.T) {
	ctx, client := setup()

	driveItems, err := client.DriveItems.List(ctx, "")
	if err != nil {
		t.Errorf("Error: %v\n", err)
		return
	}
	for _, driveItem := range driveItems.DriveItems {
		fmt.Printf("Results: %v\n", driveItem.Name)
	}
}

func TestDriveItems_GetItemsInSpecificFolder(t *testing.T) {
	// ctx, client := setup()

	// driveItems, err := client.DriveItems.List(ctx, "<<input>>")
	// if err != nil {
	// 	t.Errorf("Error: %v\n", err)
	// 	return
	// }
	// for _, driveItem := range driveItems.DriveItems {
	// 	fmt.Printf("Results: %v\n", driveItem.Name)
	// }
}

func TestDriveItems_GetMusicFolder(t *testing.T) {
	ctx, client := setup()

	musicDriveItem, err := client.DriveItems.GetSpecial(ctx, onedrive.Music)
	if err != nil {
		t.Errorf("Error: %v\n", err)
		return
	}
	fmt.Printf("Music DriveItem Name: %v\n", musicDriveItem.Name)
	fmt.Printf("Music DriveItem Id: %v\n", musicDriveItem.Id)
}

func TestDriveItems_GetItemsInMusicFolder(t *testing.T) {
	ctx, client := setup()

	musicDriveItems, err := client.DriveItems.ListSpecial(ctx, onedrive.Music)
	if err != nil {
		t.Errorf("Error: %v\n", err)
		return
	}
	for _, driveItem := range musicDriveItems.DriveItems {
		fmt.Printf("Results: %v\n", driveItem.Name)
	}
}

func TestDriveItems_CreateNewFolders(t *testing.T) {
	// ctx, client := setup()

	// newFolder, err := client.DriveItems.CreateNewFolder(ctx, "", "", "New Folder")
	// if err != nil {
	// 	t.Errorf("Error: %v\n", err)
	// 	return
	// }
	// fmt.Printf("New Folder Name: %v\n", newFolder.Name)
	// fmt.Printf("New Folder Id: %v\n", newFolder.Id)

	// // create a new subfolder "Inner SubFolder" in the "New Folder" created above for the authenticated user
	// newSubFolder, err := client.DriveItems.CreateNewFolder(ctx, "", newFolder.Id, "Inner SubFolder")
	// if err != nil {
	// 	t.Errorf("Error: %v\n", err)
	// 	return
	// }
	// fmt.Printf("New SubFolder Name: %v\n", newSubFolder.Name)
	// fmt.Printf("New SubFolder Id: %v\n", newSubFolder.Id)

	// // create a new folder "New Folder A" in the root of a selected drive for the authenticated user
	// newFolderA, err := client.DriveItems.CreateNewFolder(ctx, "<<input>>", "", "New Folder A")
	// if err != nil {
	// 	t.Errorf("Error: %v\n", err)
	// 	return
	// }
	// fmt.Printf("New Folder Name: %v\n", newFolderA.Name)
	// fmt.Printf("New Folder Id: %v\n", newFolderA.Id)
}

func TestDriveItems_Move(t *testing.T) {
	//ctx, client := setup()

	// _, err = client.DriveItems.Move(ctx, "", "<<input>>", "<<input>>")
	// if err != nil {
	// 	t.Errorf("Error: %v\n", err)
	// 	return
	// }
}

func TestDriveItems_Delete(t *testing.T) {
	// ctx, client := setup()

	// err := client.DriveItems.Delete(ctx, "", "<<input>>")
	// if err != nil {
	// 	t.Errorf("Error: %v\n", err)
	// 	return
	// }
}

func TestDriveItems_RenameItem(t *testing.T) {
	// ctx, client := setup()

	// renameResponse, err := client.DriveItems.Rename(ctx, "", "<<input>>", "Test 1.txt")
	// if err != nil {
	// 	t.Errorf("Error: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Results: %v\n", renameResponse.Name)
}

func TestDriveItems_CopyItem(t *testing.T) {
	// ctx, client := setup()

	// copyResponse, err := client.DriveItems.Copy(ctx, "", "<<input>>", "", "<<input>>", "Test 2.txt")
	// if err != nil {
	// 	t.Errorf("Error: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Location: %v\n", copyResponse.Location)
}

func TestDriveItems_CopyFolder(t *testing.T) {
	// ctx, client := setup()

	//copyResponse, err = client.DriveItems.Copy(ctx, "", "<<input>>", "", "<<input>>", "New Folder")
	// if err != nil {
	// 	t.Errorf("Error: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Location: %v\n", copyResponse.Location)
}

func TestDriveItems_UploadFile(t *testing.T) {
	// ctx, client := setup()

	// uploadedDriveItem, err := client.DriveItems.UploadNewFile(ctx, "", "<<input>>", `<<input>>`)
	// if err != nil {
	// 	t.Errorf("Error: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Uploaded DriveItem: %v\n", uploadedDriveItem)
}

func TestDriveItems_UploadFileAndReplace(t *testing.T) {
	// ctx, client := setup()

	// uploadedDriveItem, err := client.DriveItems.UploadToReplaceFile(ctx, "", `<<input>>`, "<<input>>")
	// if err != nil {
	// 	t.Errorf("Error: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Uploaded DriveItem: %v\n", uploadedDriveItem)
}
