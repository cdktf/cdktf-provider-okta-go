// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userschemaproperty


type UserSchemaPropertyMasterOverridePriority struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/user_schema_property#value UserSchemaProperty#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/user_schema_property#type UserSchemaProperty#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

