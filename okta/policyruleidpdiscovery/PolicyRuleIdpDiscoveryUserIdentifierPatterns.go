// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policyruleidpdiscovery


type PolicyRuleIdpDiscoveryUserIdentifierPatterns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.9.0/docs/resources/policy_rule_idp_discovery#match_type PolicyRuleIdpDiscovery#match_type}.
	MatchType *string `field:"optional" json:"matchType" yaml:"matchType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.9.0/docs/resources/policy_rule_idp_discovery#value PolicyRuleIdpDiscovery#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

