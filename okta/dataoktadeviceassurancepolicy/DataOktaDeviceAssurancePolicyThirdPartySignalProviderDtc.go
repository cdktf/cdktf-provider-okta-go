// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataoktadeviceassurancepolicy


type DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtc struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#allow_screen_lock DataOktaDeviceAssurancePolicy#allow_screen_lock}.
	AllowScreenLock interface{} `field:"optional" json:"allowScreenLock" yaml:"allowScreenLock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#browser_version DataOktaDeviceAssurancePolicy#browser_version}.
	BrowserVersion *DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcBrowserVersion `field:"optional" json:"browserVersion" yaml:"browserVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#built_in_dns_client_enabled DataOktaDeviceAssurancePolicy#built_in_dns_client_enabled}.
	BuiltInDnsClientEnabled interface{} `field:"optional" json:"builtInDnsClientEnabled" yaml:"builtInDnsClientEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#chrome_remote_desktop_app_blocked DataOktaDeviceAssurancePolicy#chrome_remote_desktop_app_blocked}.
	ChromeRemoteDesktopAppBlocked interface{} `field:"optional" json:"chromeRemoteDesktopAppBlocked" yaml:"chromeRemoteDesktopAppBlocked"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#crowd_strike_agent_id DataOktaDeviceAssurancePolicy#crowd_strike_agent_id}.
	CrowdStrikeAgentId *string `field:"optional" json:"crowdStrikeAgentId" yaml:"crowdStrikeAgentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#crowd_strike_customer_id DataOktaDeviceAssurancePolicy#crowd_strike_customer_id}.
	CrowdStrikeCustomerId *string `field:"optional" json:"crowdStrikeCustomerId" yaml:"crowdStrikeCustomerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#device_enrollment_domain DataOktaDeviceAssurancePolicy#device_enrollment_domain}.
	DeviceEnrollmentDomain *string `field:"optional" json:"deviceEnrollmentDomain" yaml:"deviceEnrollmentDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#disk_encrypted DataOktaDeviceAssurancePolicy#disk_encrypted}.
	DiskEncrypted interface{} `field:"optional" json:"diskEncrypted" yaml:"diskEncrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#key_trust_level DataOktaDeviceAssurancePolicy#key_trust_level}.
	KeyTrustLevel *string `field:"optional" json:"keyTrustLevel" yaml:"keyTrustLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#managed_device DataOktaDeviceAssurancePolicy#managed_device}.
	ManagedDevice interface{} `field:"optional" json:"managedDevice" yaml:"managedDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#os_firewall DataOktaDeviceAssurancePolicy#os_firewall}.
	OsFirewall interface{} `field:"optional" json:"osFirewall" yaml:"osFirewall"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#os_version DataOktaDeviceAssurancePolicy#os_version}.
	OsVersion *DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOsVersion `field:"optional" json:"osVersion" yaml:"osVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#password_protection_warning_trigger DataOktaDeviceAssurancePolicy#password_protection_warning_trigger}.
	PasswordProtectionWarningTrigger *string `field:"optional" json:"passwordProtectionWarningTrigger" yaml:"passwordProtectionWarningTrigger"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#realtime_url_check_mode DataOktaDeviceAssurancePolicy#realtime_url_check_mode}.
	RealtimeUrlCheckMode interface{} `field:"optional" json:"realtimeUrlCheckMode" yaml:"realtimeUrlCheckMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#safe_browsing_protection_level DataOktaDeviceAssurancePolicy#safe_browsing_protection_level}.
	SafeBrowsingProtectionLevel *string `field:"optional" json:"safeBrowsingProtectionLevel" yaml:"safeBrowsingProtectionLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#screen_lock_secured DataOktaDeviceAssurancePolicy#screen_lock_secured}.
	ScreenLockSecured interface{} `field:"optional" json:"screenLockSecured" yaml:"screenLockSecured"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#site_isolation_enabled DataOktaDeviceAssurancePolicy#site_isolation_enabled}.
	SiteIsolationEnabled interface{} `field:"optional" json:"siteIsolationEnabled" yaml:"siteIsolationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#third_party_blocking_enabled DataOktaDeviceAssurancePolicy#third_party_blocking_enabled}.
	ThirdPartyBlockingEnabled interface{} `field:"optional" json:"thirdPartyBlockingEnabled" yaml:"thirdPartyBlockingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#windows_machine_domain DataOktaDeviceAssurancePolicy#windows_machine_domain}.
	WindowsMachineDomain *string `field:"optional" json:"windowsMachineDomain" yaml:"windowsMachineDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.16.0/docs/data-sources/device_assurance_policy#windows_user_domain DataOktaDeviceAssurancePolicy#windows_user_domain}.
	WindowsUserDomain *string `field:"optional" json:"windowsUserDomain" yaml:"windowsUserDomain"`
}

