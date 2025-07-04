// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventhook


type EventHookHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/event_hook#key EventHook#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.20.0/docs/resources/event_hook#value EventHook#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

