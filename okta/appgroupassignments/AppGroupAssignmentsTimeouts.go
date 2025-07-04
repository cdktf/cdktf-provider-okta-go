// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appgroupassignments


type AppGroupAssignmentsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_group_assignments#create AppGroupAssignments#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_group_assignments#read AppGroupAssignments#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_group_assignments#update AppGroupAssignments#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

