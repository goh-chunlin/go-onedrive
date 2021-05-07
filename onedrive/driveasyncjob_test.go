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

func TestDriveAsyncJobService_Monitor_SuccessFile(t *testing.T) {
	client, mux, _, teardown := setup()

	defer teardown()

	mux.HandleFunc("/monitor/asyncJobSuccessFile", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		jsonData := getTestDataFromFile(t, "fake_asyncJobSuccessFile.json")

		fmt.Fprint(w, string(jsonData))
	})

	ctx := context.Background()
	gotOneDriveResponse, err := client.DriveAsyncJob.Monitor(ctx, "/test-onedrive-api/monitor/asyncJobSuccessFile")
	if err != nil {
		t.Errorf("DriveItems.Monitor returned error: %v", err)
	}

	jsonFile, err := os.Open("testdata/fake_asyncJobSuccessFile.json")

	if err != nil {
		t.Errorf("Cannot load the file data for comparison: %v", err)
	}

	defer jsonFile.Close()

	comparedToData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		t.Errorf("Cannot load the file data for comparison: %v", err)
	}

	var wantOneDriveResponse *OneDriveAsyncJobMonitorResponse
	json.Unmarshal(comparedToData, &wantOneDriveResponse)

	if !reflect.DeepEqual(gotOneDriveResponse, wantOneDriveResponse) {
		t.Errorf("Drives.Monitor returned %+v, want %+v", gotOneDriveResponse, wantOneDriveResponse)
	}

}

func TestDriveAsyncJobService_Monitor_SuccessFolder(t *testing.T) {
	client, mux, _, teardown := setup()

	defer teardown()

	mux.HandleFunc("/monitor/asyncJobSuccessFolder", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		jsonData := getTestDataFromFile(t, "fake_asyncJobSuccessFolder.json")

		fmt.Fprint(w, string(jsonData))
	})

	ctx := context.Background()
	gotOneDriveResponse, err := client.DriveAsyncJob.Monitor(ctx, "/test-onedrive-api/monitor/asyncJobSuccessFolder")
	if err != nil {
		t.Errorf("DriveItems.Monitor returned error: %v", err)
	}

	jsonFile, err := os.Open("testdata/fake_asyncJobSuccessFolder.json")

	if err != nil {
		t.Errorf("Cannot load the file data for comparison: %v", err)
	}

	defer jsonFile.Close()

	comparedToData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		t.Errorf("Cannot load the file data for comparison: %v", err)
	}

	var wantOneDriveResponse *OneDriveAsyncJobMonitorResponse
	json.Unmarshal(comparedToData, &wantOneDriveResponse)

	if !reflect.DeepEqual(gotOneDriveResponse, wantOneDriveResponse) {
		t.Errorf("Drives.Monitor returned %+v, want %+v", gotOneDriveResponse, wantOneDriveResponse)
	}

}

func TestDriveAsyncJobService_Monitor_Failed(t *testing.T) {
	client, mux, _, teardown := setup()

	defer teardown()

	mux.HandleFunc("/monitor/asyncJobFailed", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		jsonData := getTestDataFromFile(t, "fake_asyncJobFailed.json")

		fmt.Fprint(w, string(jsonData))
	})

	ctx := context.Background()
	gotOneDriveResponse, err := client.DriveAsyncJob.Monitor(ctx, "/test-onedrive-api/monitor/asyncJobFailed")
	if err != nil {
		t.Errorf("DriveItems.Monitor returned error: %v", err)
	}

	jsonFile, err := os.Open("testdata/fake_asyncJobFailed.json")

	if err != nil {
		t.Errorf("Cannot load the file data for comparison: %v", err)
	}

	defer jsonFile.Close()

	comparedToData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		t.Errorf("Cannot load the file data for comparison: %v", err)
	}

	var wantOneDriveResponse *OneDriveAsyncJobMonitorResponse
	json.Unmarshal(comparedToData, &wantOneDriveResponse)

	if !reflect.DeepEqual(gotOneDriveResponse, wantOneDriveResponse) {
		t.Errorf("Drives.Monitor returned %+v, want %+v", gotOneDriveResponse, wantOneDriveResponse)
	}

}
