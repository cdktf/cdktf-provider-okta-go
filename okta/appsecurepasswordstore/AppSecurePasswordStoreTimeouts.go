// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsecurepasswordstore


type AppSecurePasswordStoreTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_secure_password_store#create AppSecurePasswordStore#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_secure_password_store#read AppSecurePasswordStore#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/app_secure_password_store#update AppSecurePasswordStore#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

