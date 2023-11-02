// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policydeviceassuranceios

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicyDeviceAssuranceIosConfig struct {
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
	// Policy device assurance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.6.0/docs/resources/policy_device_assurance_ios#name PolicyDeviceAssuranceIos#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The device jailbreak. Only for android and iOS platform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.6.0/docs/resources/policy_device_assurance_ios#jailbreak PolicyDeviceAssuranceIos#jailbreak}
	Jailbreak interface{} `field:"optional" json:"jailbreak" yaml:"jailbreak"`
	// The device os minimum version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.6.0/docs/resources/policy_device_assurance_ios#os_version PolicyDeviceAssuranceIos#os_version}
	OsVersion *string `field:"optional" json:"osVersion" yaml:"osVersion"`
	// List of screenlock type, can be BIOMETRIC or BIOMETRIC, PASSCODE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.6.0/docs/resources/policy_device_assurance_ios#screenlock_type PolicyDeviceAssuranceIos#screenlock_type}
	ScreenlockType *[]*string `field:"optional" json:"screenlockType" yaml:"screenlockType"`
}

