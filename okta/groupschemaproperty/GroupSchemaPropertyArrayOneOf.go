// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupschemaproperty


type GroupSchemaPropertyArrayOneOf struct {
	// Value mapping to member of `array_enum`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/group_schema_property#const GroupSchemaProperty#const}
	Const *string `field:"required" json:"const" yaml:"const"`
	// Display name for the enum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/group_schema_property#title GroupSchemaProperty#title}
	Title *string `field:"required" json:"title" yaml:"title"`
}

