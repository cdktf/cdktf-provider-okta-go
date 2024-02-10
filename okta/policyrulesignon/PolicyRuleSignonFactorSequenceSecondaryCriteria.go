// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policyrulesignon


type PolicyRuleSignonFactorSequenceSecondaryCriteria struct {
	// Type of a Factor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.7.0/docs/resources/policy_rule_signon#factor_type PolicyRuleSignon#factor_type}
	FactorType *string `field:"required" json:"factorType" yaml:"factorType"`
	// Factor provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.7.0/docs/resources/policy_rule_signon#provider PolicyRuleSignon#provider}
	Provider *string `field:"required" json:"provider" yaml:"provider"`
}

