package frontend

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/ugorji/go/codec"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database/cosmosdb"
	"github.com/Azure/ARO-RP/pkg/frontend/middleware"
)

func (f *frontend) getAsyncOperation(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(middleware.ContextKeyLog).(*logrus.Entry)

	b, err := f._getAsyncOperation(r)

	reply(log, w, nil, b, err)
}

func (f *frontend) _getAsyncOperation(r *http.Request) ([]byte, error) {
	vars := mux.Vars(r)

	asyncdoc, err := f.db.AsyncOperations.Get(vars["operationId"])
	switch {
	case cosmosdb.IsErrorStatusCode(err, http.StatusNotFound):
		return nil, api.NewCloudError(http.StatusNotFound, api.CloudErrorCodeNotFound, "", "The entity was not found.")
	case err != nil:
		return nil, err
	}

	doc, err := f.db.OpenShiftClusters.Get(asyncdoc.OpenShiftClusterKey)
	if err != nil && !cosmosdb.IsErrorStatusCode(err, http.StatusNotFound) {
		return nil, err
	}

	// don't give away the final operation status until it's committed to the
	// database
	if doc != nil && doc.AsyncOperationID == vars["operationId"] {
		asyncdoc.AsyncOperation.ProvisioningState = asyncdoc.AsyncOperation.InitialProvisioningState
		asyncdoc.AsyncOperation.EndTime = nil
		asyncdoc.AsyncOperation.Error = nil
	}

	asyncdoc.AsyncOperation.MissingFields = api.MissingFields{}
	asyncdoc.AsyncOperation.InitialProvisioningState = ""

	h := &codec.JsonHandle{
		Indent: 4,
	}

	var b []byte
	err = codec.NewEncoderBytes(&b, h).Encode(asyncdoc.AsyncOperation)
	if err != nil {
		return nil, err
	}

	return b, nil
}