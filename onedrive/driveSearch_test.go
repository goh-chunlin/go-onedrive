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

func TestDriveSearchService_SearchWithEmptyQuery_authenticatedUser(t *testing.T) {
	client, mux, _, teardown := setup()

	defer teardown()
	mux.HandleFunc("/me/drive/root/search(q='')", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		jsonData := getTestDataFromFile(t, "fake_driveItems_withEmptyQuery_searchResults.json")

		fmt.Fprint(w, string(jsonData))
	})

	ctx := context.Background()
	_, err := client.DriveSearch.Search(ctx, "")
	if err == nil {
		t.Errorf("There should be an error")
	}

}

func TestDriveSearchService_SearchDriveItems_authenticatedUser(t *testing.T) {
	client, mux, _, teardown := setup()

	defer teardown()
	mux.HandleFunc("/me/drive/root/search(q='Test')", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		jsonData := getTestDataFromFile(t, "fake_driveItems_searchResults.json")

		fmt.Fprint(w, string(jsonData))
	})

	ctx := context.Background()
	gotOneDriveResponse, err := client.DriveSearch.Search(ctx, "Test")
	if err != nil {
		t.Errorf("DriveSearch.Search returned error: %v", err)
	}

	jsonFile, err := os.Open("testdata/fake_driveItems_searchResults.json")

	if err != nil {
		t.Errorf("Cannot load the file data for comparison: %v", err)
	}

	defer jsonFile.Close()

	comparedToData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		t.Errorf("Cannot load the file data for comparison: %v", err)
	}

	var wantOneDriveResponse *OneDriveDriveSearchResponse
	json.Unmarshal(comparedToData, &wantOneDriveResponse)

	if !reflect.DeepEqual(gotOneDriveResponse, wantOneDriveResponse) {
		t.Errorf("DriveSearch.Search returned %+v, want %+v", gotOneDriveResponse, wantOneDriveResponse)
	}

}
