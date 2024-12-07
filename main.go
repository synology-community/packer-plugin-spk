// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"

	"github.com/synology-community/packer-plugin-synopkg/builder/synopkg"
	synopkgData "github.com/synology-community/packer-plugin-synopkg/datasource/synopkg"
	synopkgPP "github.com/synology-community/packer-plugin-synopkg/post-processor/synopkg"
	synopkgProv "github.com/synology-community/packer-plugin-synopkg/provisioner/synopkg"
	synopkgVersion "github.com/synology-community/packer-plugin-synopkg/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder("my-builder", new(synopkg.Builder))
	pps.RegisterProvisioner("my-provisioner", new(synopkgProv.Provisioner))
	pps.RegisterPostProcessor("my-post-processor", new(synopkgPP.PostProcessor))
	pps.RegisterDatasource("my-datasource", new(synopkgData.Datasource))
	pps.SetVersion(synopkgVersion.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
