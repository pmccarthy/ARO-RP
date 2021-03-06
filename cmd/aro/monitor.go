package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/database"
	"github.com/Azure/ARO-RP/pkg/env"
	"github.com/Azure/ARO-RP/pkg/metrics/statsd"
	pkgmonitor "github.com/Azure/ARO-RP/pkg/monitor"
	"github.com/Azure/ARO-RP/pkg/util/encryption"
)

func monitor(ctx context.Context, log *logrus.Entry) error {
	uuid := uuid.NewV4().String()
	log.Printf("uuid %s", uuid)

	_env, err := env.NewEnv(ctx, log)
	if err != nil {
		return err
	}

	if _, ok := _env.(env.Dev); !ok {
		for _, key := range []string{
			"MDM_ACCOUNT",
			"MDM_NAMESPACE",
		} {
			if _, found := os.LookupEnv(key); !found {
				return fmt.Errorf("environment variable %q unset", key)
			}
		}
	}

	m, err := statsd.New(ctx, log.WithField("component", "metrics"), _env, os.Getenv("MDM_ACCOUNT"), os.Getenv("MDM_NAMESPACE"))
	if err != nil {
		return err
	}

	clusterm, err := statsd.New(ctx, log.WithField("component", "metrics"), _env, os.Getenv("CLUSTER_MDM_ACCOUNT"), os.Getenv("CLUSTER_MDM_NAMESPACE"))
	if err != nil {
		return err
	}

	cipher, err := encryption.NewXChaCha20Poly1305(ctx, _env)
	if err != nil {
		return err
	}

	db, err := database.NewDatabase(ctx, log.WithField("component", "database"), _env, m, cipher, uuid)
	if err != nil {
		return err
	}

	mon := pkgmonitor.NewMonitor(log.WithField("component", "monitor"), _env, db, m, clusterm)

	return mon.Run(ctx)
}
