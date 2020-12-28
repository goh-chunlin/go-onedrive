// Copyright 2020 The go-onedrive AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package onedrive

// ErrorResponse represents the error response returned by OneDrive drive API.
type ErrorResponse struct {
	Error *Error `json:"error"`
}

// Error represents the error in the response returned by OneDrive drive API.
type Error struct {
	Code             string      `json:"code"`
	Message          string      `json:"message"`
	LocalizedMessage string      `json:"localizedMessage"`
	InnerError       *InnerError `json:"innerError"`
}

// InnerError represents the error details in the error returned by OneDrive drive API.
type InnerError struct {
	Date            string `json:"date"`
	RequestId       string `json:"request-id"`
	ClientRequestId string `json:"client-request-id"`
}
