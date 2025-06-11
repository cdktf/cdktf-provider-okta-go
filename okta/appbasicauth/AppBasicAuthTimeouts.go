// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appbasicauth


type AppBasicAuthTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_basic_auth#create AppBasicAuth#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_basic_auth#read AppBasicAuth#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_basic_auth#update AppBasicAuth#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

