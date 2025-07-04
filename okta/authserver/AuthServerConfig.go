// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authserver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthServerConfig struct {
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
	// The recipients that the tokens are intended for.
	//
	// This becomes the `aud` claim in an access token. Currently Okta only supports a single value here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/auth_server#audiences AuthServer#audiences}
	Audiences *[]*string `field:"required" json:"audiences" yaml:"audiences"`
	// The name of the authorization server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/auth_server#name AuthServer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The key rotation mode for the authorization server. Can be `AUTO` or `MANUAL`. Default: `AUTO`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/auth_server#credentials_rotation_mode AuthServer#credentials_rotation_mode}
	CredentialsRotationMode *string `field:"optional" json:"credentialsRotationMode" yaml:"credentialsRotationMode"`
	// The description of the authorization server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/auth_server#description AuthServer#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/auth_server#id AuthServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// *Early Access Property*.
	//
	// Allows you to use a custom issuer URL. It can be set to `CUSTOM_URL`, `ORG_URL`, or `DYNAMIC`. Default: `ORG_URL`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/auth_server#issuer_mode AuthServer#issuer_mode}
	IssuerMode *string `field:"optional" json:"issuerMode" yaml:"issuerMode"`
	// Default to `ACTIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/auth_server#status AuthServer#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

