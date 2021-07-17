// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package integration

import (
	"testing"
)

func TestPermission_CreateAnynomousViewLink(t *testing.T) {
	// ctx, client := setup()

	// permission, err := client.DrivePermissions.CreateShareLink(ctx, "<<input>>", onedrive.View, onedrive.Anonymous)
	// if err != nil {
	// 	t.Errorf("Error: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Result: %v\n", permission.Link.Type+"; "+permission.Link.Scope+"; "+permission.Link.URL)
}

func TestPermission_List(t *testing.T) {
	// ctx, client := setup()

	// permissions, err := client.DrivePermissions.List(ctx, "<<input>>")
	// if err != nil {
	// 	t.Errorf("Error: %v\n", err)
	// 	return
	// }

	// fmt.Printf("Number of Permissions: %v\n", len(permissions))
	// for _, permission := range permissions {
	// 	fmt.Printf("Results: %v\n", permission.ID+"; "+permission.Link.Type+"; "+permission.Link.Scope+"; "+permission.Link.URL)
	// }
}

func TestPermission_Delete(t *testing.T) {
	// ctx, client := setup()

	// err := client.DrivePermissions.Delete(ctx, "", "<<input>>", "<<input>>")
	// if err != nil {
	// 	t.Errorf("Error: %v\n", err)
	// 	return
	// }
}
