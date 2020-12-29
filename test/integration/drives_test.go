package integration

import (
	"fmt"
	"testing"
)

func TestDrives_GetDefaultDrive(t *testing.T) {
	ctx, client := setup()

	defaultDrive, err := client.Drives.Get(ctx, "")
	if err != nil {
		t.Errorf("Error: %v\n", err)
		return
	}
	fmt.Printf("Default Drive ID: %v\n", defaultDrive.Id)
	fmt.Printf("Default Drive Owner: %v\n", defaultDrive.Owner.User.DisplayName)
}

func TestDrives_GetDriveById(t *testing.T) {
	// ctx, client := setup()

	// specifiedDrive, err := client.Drives.Get(ctx, "<<input>>")
	// if err != nil {
	// 	t.Errorf("Error: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Specified Drive ID: %v\n", specifiedDrive.Id)
	// fmt.Printf("Specified Drive Owner: %v\n", specifiedDrive.Owner.User.DisplayName)
}

func TestDrives_GetAllDrives(t *testing.T) {
	ctx, client := setup()

	drives, err := client.Drives.List(ctx)
	if err != nil {
		t.Errorf("Error: %v\n", err)
		return
	}
	for _, drive := range drives.Drives {
		fmt.Printf("Results: %v\n", drive.Owner.User.DisplayName)
	}
}
