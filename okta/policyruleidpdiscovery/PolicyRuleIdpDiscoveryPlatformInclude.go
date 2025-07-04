// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policyruleidpdiscovery


type PolicyRuleIdpDiscoveryPlatformInclude struct {
	// Only available with OTHER OS type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/policy_rule_idp_discovery#os_expression PolicyRuleIdpDiscovery#os_expression}
	OsExpression *string `field:"optional" json:"osExpression" yaml:"osExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/policy_rule_idp_discovery#os_type PolicyRuleIdpDiscovery#os_type}.
	OsType *string `field:"optional" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/policy_rule_idp_discovery#type PolicyRuleIdpDiscovery#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

