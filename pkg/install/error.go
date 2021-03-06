package install

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// isResourceQuotaExceededError returns true if it is Quota error with the
// original error message
func isResourceQuotaExceededError(err error) (bool, string) {
	if detailedErr, ok := err.(autorest.DetailedError); ok {
		// error format:
		// (autorest.DetailedError).Original.(*azure.ServiceError).Details.([]map[string]interface{})
		if serviceErr, ok := detailedErr.Original.(*azure.ServiceError); ok {
			for _, d := range serviceErr.Details {
				if code, ok := d["code"].(string); ok && code == "QuotaExceeded" {
					if message, ok := d["message"].(string); ok {
						return true, message
					}
				}
			}
		}
	}
	return false, ""
}

// isDeploymentActiveError returns true it is deployment active error
func isDeploymentActiveError(err error) bool {
	if detailedErr, ok := err.(autorest.DetailedError); ok {
		if requestErr, ok := detailedErr.Original.(azure.RequestError); ok &&
			requestErr.ServiceError != nil &&
			requestErr.ServiceError.Code == "DeploymentActive" {
			return true
		}
	}
	return false
}
