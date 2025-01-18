// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataoktaapps

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOktaAppsConfig struct {
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
	// Search only active applications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.13.0/docs/data-sources/apps#active_only DataOktaApps#active_only}
	ActiveOnly interface{} `field:"optional" json:"activeOnly" yaml:"activeOnly"`
	// Specifies whether to include non-active, but not deleted apps in the results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.13.0/docs/data-sources/apps#include_non_deleted DataOktaApps#include_non_deleted}
	IncludeNonDeleted interface{} `field:"optional" json:"includeNonDeleted" yaml:"includeNonDeleted"`
	// Searches for applications whose label or name property matches this value exactly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.13.0/docs/data-sources/apps#label DataOktaApps#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Searches for applications whose label or name property begins with this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.13.0/docs/data-sources/apps#label_prefix DataOktaApps#label_prefix}
	LabelPrefix *string `field:"optional" json:"labelPrefix" yaml:"labelPrefix"`
	// Specifies whether to use query optimization.
	//
	// If you specify `useOptimization=true` in the request query, the response contains a subset of app instance properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.13.0/docs/data-sources/apps#use_optimization DataOktaApps#use_optimization}
	UseOptimization interface{} `field:"optional" json:"useOptimization" yaml:"useOptimization"`
}

