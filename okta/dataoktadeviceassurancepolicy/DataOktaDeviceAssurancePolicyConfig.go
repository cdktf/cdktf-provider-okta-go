// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataoktadeviceassurancepolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOktaDeviceAssurancePolicyConfig struct {
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
	// ID of the user type to retrieve, conflicts with `name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.14.0/docs/data-sources/device_assurance_policy#id DataOktaDeviceAssurancePolicy#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name of user type to retrieve, conflicts with `id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.14.0/docs/data-sources/device_assurance_policy#name DataOktaDeviceAssurancePolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicates if the device contains a secure hardware functionality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.14.0/docs/data-sources/device_assurance_policy#secure_hardware_present DataOktaDeviceAssurancePolicy#secure_hardware_present}
	SecureHardwarePresent interface{} `field:"optional" json:"secureHardwarePresent" yaml:"secureHardwarePresent"`
	// Indicates if the device contains a secure hardware functionality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.14.0/docs/data-sources/device_assurance_policy#third_party_signal_provider DataOktaDeviceAssurancePolicy#third_party_signal_provider}
	ThirdPartySignalProvider *DataOktaDeviceAssurancePolicyThirdPartySignalProvider `field:"optional" json:"thirdPartySignalProvider" yaml:"thirdPartySignalProvider"`
}

