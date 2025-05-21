// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emailsmtpserver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmailSmtpServerConfig struct {
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
	// Human-readable name for your SMTP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.19.0/docs/resources/email_smtp_server#alias EmailSmtpServer#alias}
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// Hostname or IP address of your SMTP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.19.0/docs/resources/email_smtp_server#host EmailSmtpServer#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// User name of the email domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.19.0/docs/resources/email_smtp_server#password EmailSmtpServer#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Port number of your SMTP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.19.0/docs/resources/email_smtp_server#port EmailSmtpServer#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Display name of the email domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.19.0/docs/resources/email_smtp_server#username EmailSmtpServer#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// If true, routes all email traffic through your SMTP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.19.0/docs/resources/email_smtp_server#enabled EmailSmtpServer#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.19.0/docs/resources/email_smtp_server#id EmailSmtpServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

