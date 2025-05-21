// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appswa


type AppSwaTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.19.0/docs/resources/app_swa#create AppSwa#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.19.0/docs/resources/app_swa#read AppSwa#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.19.0/docs/resources/app_swa#update AppSwa#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

