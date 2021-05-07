package onedrive

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateSharingLink(t *testing.T) {
	client, mux, _, teardown := setup()

	defer teardown()

	jsonData := getTestDataFromFile(t, "fake_permission.json")
	mux.HandleFunc("/me/drive/items/1/createLink", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		fmt.Fprint(w, string(jsonData))
	})

	ctx := context.Background()
	gotOneDriveResponse, err := client.DrivePermissions.CreateShareLink(ctx, "1", "read", "anonymous")
	if err != nil {
		t.Errorf("CreateShareLink returned error: %v", err)
	}

	var wantDriveItem *Permission
	if err := json.Unmarshal(jsonData, &wantDriveItem); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(gotOneDriveResponse, wantDriveItem) {
		t.Errorf("CreateShareLink returned %+v, want %+v", gotOneDriveResponse, wantDriveItem)
	}
}

func TestListPermissions(t *testing.T) {
	client, mux, _, teardown := setup()

	defer teardown()

	jsonData := getTestDataFromFile(t, "fake_permissions.json")
	mux.HandleFunc("/me/drive/items/1/permissions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		fmt.Fprint(w, string(jsonData))
	})

	ctx := context.Background()
	gotOneDriveResponse, err := client.DrivePermissions.List(ctx, "1")
	if err != nil {
		t.Errorf("List returned error: %v", err)
	}

	var wantDriveItem *ListPermissionsResponse
	if err := json.Unmarshal(jsonData, &wantDriveItem); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(gotOneDriveResponse, wantDriveItem.Value) {
		t.Errorf("List returned %+v, want %+v", gotOneDriveResponse, wantDriveItem)
	}
}
