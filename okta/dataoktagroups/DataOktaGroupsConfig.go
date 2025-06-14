// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataoktagroups

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOktaGroupsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/data-sources/groups#id DataOktaGroups#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The maximum number of groups returned by the Okta API, between 1 and 10000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/data-sources/groups#limit DataOktaGroups#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Searches the name property of groups for matching value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/data-sources/groups#q DataOktaGroups#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// Searches for groups with a supported filtering expression for all attributes except for '_embedded', '_links', and 'objectClass'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/data-sources/groups#search DataOktaGroups#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Type of the group.
	//
	// When specified in the terraform resource, will act as a filter when searching for the groups
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/data-sources/groups#type DataOktaGroups#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

