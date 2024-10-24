// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package apigee

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/tampakrap/provider-upjet-gcp-beta/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_apigee_endpoint_attachment", func(r *config.Resource) {
		r.References["service_attachment"] = config.Reference{
			TerraformName: "google_compute_service_attachment",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})
}
