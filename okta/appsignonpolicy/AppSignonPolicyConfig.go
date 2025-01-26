// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsignonpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppSignonPolicyConfig struct {
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
	// Description of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.13.1/docs/resources/app_signon_policy#description AppSignonPolicy#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Name of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.13.1/docs/resources/app_signon_policy#name AppSignonPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Default rules of the policy set to `DENY` or not.
	//
	// If `false`, it is set to `DENY`. **WARNING** setting this attribute to false change the OKTA default behavior. Use at your own risk. This is only apply during creation, so import or update will not work
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.13.1/docs/resources/app_signon_policy#catch_all AppSignonPolicy#catch_all}
	CatchAll interface{} `field:"optional" json:"catchAll" yaml:"catchAll"`
}

