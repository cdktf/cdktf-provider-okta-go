// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appgroupassignment


type AppGroupAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_group_assignment#create AppGroupAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_group_assignment#read AppGroupAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_group_assignment#update AppGroupAssignment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

