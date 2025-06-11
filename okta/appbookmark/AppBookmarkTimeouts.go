// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appbookmark


type AppBookmarkTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_bookmark#create AppBookmark#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_bookmark#read AppBookmark#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_bookmark#update AppBookmark#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

