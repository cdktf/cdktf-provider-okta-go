// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appoauth


type AppOauthTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_oauth#create AppOauth#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_oauth#read AppOauth#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_oauth#update AppOauth#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

