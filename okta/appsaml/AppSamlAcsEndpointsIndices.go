// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsaml


type AppSamlAcsEndpointsIndices struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.19.0/docs/resources/app_saml#index AppSaml#index}.
	Index *float64 `field:"required" json:"index" yaml:"index"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.19.0/docs/resources/app_saml#url AppSaml#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

