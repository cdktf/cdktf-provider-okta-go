// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupschemaproperty


type GroupSchemaPropertyMasterOverridePriority struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/group_schema_property#value GroupSchemaProperty#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/group_schema_property#type GroupSchemaProperty#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

