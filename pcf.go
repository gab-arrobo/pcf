// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * Npcf_BDTPolicyControl Service API
 *
 * The Npcf_BDTPolicyControl Service is used by an NF service consumer to
 * retrieve background data transfer policies from the PCF and to update the PCF with
 * the background data transfer policy selected by the NF service consumer.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/omec-project/pcf/logger"
	"github.com/omec-project/pcf/service"
	"github.com/urfave/cli/v3"
)

var PCF = &service.PCF{}

func main() {
	app := &cli.Command{}
	app.Name = "pcf"
	logger.AppLog.Infoln(app.Name)
	app.Usage = "Policy Control Function"
	app.UsageText = "pcf -cfg <pcf_config_file.conf>"
	app.Action = action
	app.Flags = PCF.GetCliCmd()

	if err := app.Run(context.Background(), os.Args); err != nil {
		logger.AppLog.Fatalf("PCF run error: %v", err)
	}
}

func action(ctx context.Context, c *cli.Command) error {
	if err := PCF.Initialize(c); err != nil {
		logger.CfgLog.Errorf("%+v", err)
		return fmt.Errorf("failed to initialize")
	}

	PCF.Start()

	return nil
}
