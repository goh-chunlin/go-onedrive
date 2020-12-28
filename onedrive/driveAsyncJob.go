// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package onedrive

import "context"

// DriveAsyncJobService handles communication with the drive items searching related methods of the OneDrive API.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/concepts/long-running-actions?view=odsp-graph-online
type DriveAsyncJobService service

// OneDriveAsyncJobMonitorResponse represents the JSON object returned by the OneDrive Async Job Monitoring API.
type OneDriveAsyncJobMonitorResponse struct {
	ErrorCode           string  `json:"errorCode"`
	ResourceId          string  `json:"resourceId"`
	Operation           string  `json:"operation"`
	Status              string  `json:"status"`
	StatusDescription   string  `json:"statusDescription"`
	PercentageCompleted float64 `json:"percentageCompleted"`
}

// Retrieve a status report from the monitor URL of OneDrive.
//
// OneDrive API docs: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/concepts/long-running-actions?view=odsp-graph-online#retrieve-a-status-report-from-the-monitor-url
func (s *DriveAsyncJobService) Monitor(ctx context.Context, monitorUrl string) (*OneDriveAsyncJobMonitorResponse, error) {
	req, err := s.client.NewRequestToOneDrive("GET", monitorUrl, nil)
	if err != nil {
		return nil, err
	}

	var oneDriveResponse *OneDriveAsyncJobMonitorResponse
	err = s.client.Do(ctx, req, &oneDriveResponse)
	if err != nil {
		return nil, err
	}

	return oneDriveResponse, nil
}
