// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package onedrive

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestDriveItemsService_ListRoot_authenticatedUser(t *testing.T) {
	client, mux, _, teardown := setup()

	defer teardown()

	mux.HandleFunc("/me/drive/root/children", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		jsonData := getTestDataFromFile(t, "fake_driveItems.json")

		fmt.Fprint(w, string(jsonData))
	})

	ctx := context.Background()
	gotOneDriveResponse, err := client.DriveItems.List(ctx, "")
	if err != nil {
		t.Errorf("DriveItems.List returned error: %v", err)
	}

	jsonFile, err := os.Open("testdata/fake_driveItems.json")

	if err != nil {
		t.Errorf("Cannot load the file data for comparison: %v", err)
	}

	defer jsonFile.Close()

	comparedToData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		t.Errorf("Cannot load the file data for comparison: %v", err)
	}

	var wantOneDriveResponse *OneDriveDriveItemsResponse
	json.Unmarshal(comparedToData, &wantOneDriveResponse)

	if !reflect.DeepEqual(gotOneDriveResponse, wantOneDriveResponse) {
		t.Errorf("Drives.List returned %+v, want %+v", gotOneDriveResponse, wantOneDriveResponse)
	}

}

func TestDriveItemsService_Get_authenticatedUser(t *testing.T) {
	client, mux, _, teardown := setup()

	defer teardown()

	mux.HandleFunc("/me/drive/items/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		jsonData := getTestDataFromFile(t, "fake_driveItem.json")

		fmt.Fprint(w, string(jsonData))
	})

	ctx := context.Background()
	gotDriveItem, err := client.DriveItems.Get(ctx, "1")
	if err != nil {
		t.Errorf("DriveItems.Get returned error: %v", err)
	}

	jsonFile, err := os.Open("testdata/fake_driveItem.json")

	if err != nil {
		t.Errorf("Cannot load the file data for comparison: %v", err)
	}

	defer jsonFile.Close()

	comparedToData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		t.Errorf("Cannot load the file data for comparison: %v", err)
	}

	var wantDriveItem *DriveItem
	json.Unmarshal(comparedToData, &wantDriveItem)

	if !reflect.DeepEqual(gotDriveItem, wantDriveItem) {
		t.Errorf("Drives.Item returned %+v, want %+v", gotDriveItem, wantDriveItem)
	}

}
