// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataoktadeviceassurancepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-okta-go/okta/v13/jsii"

	"github.com/cdktf/cdktf-provider-okta-go/okta/v13/dataoktadeviceassurancepolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference interface {
	cdktf.ComplexObject
	AllowScreenLock() interface{}
	SetAllowScreenLock(val interface{})
	AllowScreenLockInput() interface{}
	BrowserVersion() DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcBrowserVersionOutputReference
	BrowserVersionInput() interface{}
	BuiltInDnsClientEnabled() interface{}
	SetBuiltInDnsClientEnabled(val interface{})
	BuiltInDnsClientEnabledInput() interface{}
	ChromeRemoteDesktopAppBlocked() interface{}
	SetChromeRemoteDesktopAppBlocked(val interface{})
	ChromeRemoteDesktopAppBlockedInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CrowdStrikeAgentId() *string
	SetCrowdStrikeAgentId(val *string)
	CrowdStrikeAgentIdInput() *string
	CrowdStrikeCustomerId() *string
	SetCrowdStrikeCustomerId(val *string)
	CrowdStrikeCustomerIdInput() *string
	DeviceEnrollmentDomain() *string
	SetDeviceEnrollmentDomain(val *string)
	DeviceEnrollmentDomainInput() *string
	DiskEncrypted() interface{}
	SetDiskEncrypted(val interface{})
	DiskEncryptedInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyTrustLevel() *string
	SetKeyTrustLevel(val *string)
	KeyTrustLevelInput() *string
	ManagedDevice() interface{}
	SetManagedDevice(val interface{})
	ManagedDeviceInput() interface{}
	OsFirewall() interface{}
	SetOsFirewall(val interface{})
	OsFirewallInput() interface{}
	OsVersion() DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOsVersionOutputReference
	OsVersionInput() interface{}
	PasswordProtectionWarningTrigger() *string
	SetPasswordProtectionWarningTrigger(val *string)
	PasswordProtectionWarningTriggerInput() *string
	RealtimeUrlCheckMode() interface{}
	SetRealtimeUrlCheckMode(val interface{})
	RealtimeUrlCheckModeInput() interface{}
	SafeBrowsingProtectionLevel() *string
	SetSafeBrowsingProtectionLevel(val *string)
	SafeBrowsingProtectionLevelInput() *string
	ScreenLockSecured() interface{}
	SetScreenLockSecured(val interface{})
	ScreenLockSecuredInput() interface{}
	SiteIsolationEnabled() interface{}
	SetSiteIsolationEnabled(val interface{})
	SiteIsolationEnabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThirdPartyBlockingEnabled() interface{}
	SetThirdPartyBlockingEnabled(val interface{})
	ThirdPartyBlockingEnabledInput() interface{}
	WindowsMachineDomain() *string
	SetWindowsMachineDomain(val *string)
	WindowsMachineDomainInput() *string
	WindowsUserDomain() *string
	SetWindowsUserDomain(val *string)
	WindowsUserDomainInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutBrowserVersion(value *DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcBrowserVersion)
	PutOsVersion(value *DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOsVersion)
	ResetAllowScreenLock()
	ResetBrowserVersion()
	ResetBuiltInDnsClientEnabled()
	ResetChromeRemoteDesktopAppBlocked()
	ResetCrowdStrikeAgentId()
	ResetCrowdStrikeCustomerId()
	ResetDeviceEnrollmentDomain()
	ResetDiskEncrypted()
	ResetKeyTrustLevel()
	ResetManagedDevice()
	ResetOsFirewall()
	ResetOsVersion()
	ResetPasswordProtectionWarningTrigger()
	ResetRealtimeUrlCheckMode()
	ResetSafeBrowsingProtectionLevel()
	ResetScreenLockSecured()
	ResetSiteIsolationEnabled()
	ResetThirdPartyBlockingEnabled()
	ResetWindowsMachineDomain()
	ResetWindowsUserDomain()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference
type jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) AllowScreenLock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowScreenLock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) AllowScreenLockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowScreenLockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) BrowserVersion() DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcBrowserVersionOutputReference {
	var returns DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcBrowserVersionOutputReference
	_jsii_.Get(
		j,
		"browserVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) BrowserVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browserVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) BuiltInDnsClientEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"builtInDnsClientEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) BuiltInDnsClientEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"builtInDnsClientEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ChromeRemoteDesktopAppBlocked() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chromeRemoteDesktopAppBlocked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ChromeRemoteDesktopAppBlockedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chromeRemoteDesktopAppBlockedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) CrowdStrikeAgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crowdStrikeAgentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) CrowdStrikeAgentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crowdStrikeAgentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) CrowdStrikeCustomerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crowdStrikeCustomerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) CrowdStrikeCustomerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crowdStrikeCustomerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) DeviceEnrollmentDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceEnrollmentDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) DeviceEnrollmentDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceEnrollmentDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) DiskEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) DiskEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) KeyTrustLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyTrustLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) KeyTrustLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyTrustLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ManagedDevice() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ManagedDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) OsFirewall() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"osFirewall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) OsFirewallInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"osFirewallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) OsVersion() DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOsVersionOutputReference {
	var returns DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOsVersionOutputReference
	_jsii_.Get(
		j,
		"osVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) OsVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"osVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) PasswordProtectionWarningTrigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordProtectionWarningTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) PasswordProtectionWarningTriggerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordProtectionWarningTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) RealtimeUrlCheckMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"realtimeUrlCheckMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) RealtimeUrlCheckModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"realtimeUrlCheckModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) SafeBrowsingProtectionLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"safeBrowsingProtectionLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) SafeBrowsingProtectionLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"safeBrowsingProtectionLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ScreenLockSecured() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"screenLockSecured",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ScreenLockSecuredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"screenLockSecuredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) SiteIsolationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"siteIsolationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) SiteIsolationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"siteIsolationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ThirdPartyBlockingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thirdPartyBlockingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ThirdPartyBlockingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thirdPartyBlockingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) WindowsMachineDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsMachineDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) WindowsMachineDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsMachineDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) WindowsUserDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsUserDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) WindowsUserDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsUserDomainInput",
		&returns,
	)
	return returns
}


func NewDataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference {
	_init_.Initialize()

	if err := validateNewDataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference_Override(d DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetAllowScreenLock(val interface{}) {
	if err := j.validateSetAllowScreenLockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowScreenLock",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetBuiltInDnsClientEnabled(val interface{}) {
	if err := j.validateSetBuiltInDnsClientEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"builtInDnsClientEnabled",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetChromeRemoteDesktopAppBlocked(val interface{}) {
	if err := j.validateSetChromeRemoteDesktopAppBlockedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chromeRemoteDesktopAppBlocked",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetCrowdStrikeAgentId(val *string) {
	if err := j.validateSetCrowdStrikeAgentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crowdStrikeAgentId",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetCrowdStrikeCustomerId(val *string) {
	if err := j.validateSetCrowdStrikeCustomerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crowdStrikeCustomerId",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetDeviceEnrollmentDomain(val *string) {
	if err := j.validateSetDeviceEnrollmentDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceEnrollmentDomain",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetDiskEncrypted(val interface{}) {
	if err := j.validateSetDiskEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncrypted",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetKeyTrustLevel(val *string) {
	if err := j.validateSetKeyTrustLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyTrustLevel",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetManagedDevice(val interface{}) {
	if err := j.validateSetManagedDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedDevice",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetOsFirewall(val interface{}) {
	if err := j.validateSetOsFirewallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osFirewall",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetPasswordProtectionWarningTrigger(val *string) {
	if err := j.validateSetPasswordProtectionWarningTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordProtectionWarningTrigger",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetRealtimeUrlCheckMode(val interface{}) {
	if err := j.validateSetRealtimeUrlCheckModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realtimeUrlCheckMode",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetSafeBrowsingProtectionLevel(val *string) {
	if err := j.validateSetSafeBrowsingProtectionLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"safeBrowsingProtectionLevel",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetScreenLockSecured(val interface{}) {
	if err := j.validateSetScreenLockSecuredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"screenLockSecured",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetSiteIsolationEnabled(val interface{}) {
	if err := j.validateSetSiteIsolationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"siteIsolationEnabled",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetThirdPartyBlockingEnabled(val interface{}) {
	if err := j.validateSetThirdPartyBlockingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thirdPartyBlockingEnabled",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetWindowsMachineDomain(val *string) {
	if err := j.validateSetWindowsMachineDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowsMachineDomain",
		val,
	)
}

func (j *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)SetWindowsUserDomain(val *string) {
	if err := j.validateSetWindowsUserDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowsUserDomain",
		val,
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) PutBrowserVersion(value *DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcBrowserVersion) {
	if err := d.validatePutBrowserVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBrowserVersion",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) PutOsVersion(value *DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOsVersion) {
	if err := d.validatePutOsVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOsVersion",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetAllowScreenLock() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowScreenLock",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetBrowserVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetBrowserVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetBuiltInDnsClientEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetBuiltInDnsClientEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetChromeRemoteDesktopAppBlocked() {
	_jsii_.InvokeVoid(
		d,
		"resetChromeRemoteDesktopAppBlocked",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetCrowdStrikeAgentId() {
	_jsii_.InvokeVoid(
		d,
		"resetCrowdStrikeAgentId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetCrowdStrikeCustomerId() {
	_jsii_.InvokeVoid(
		d,
		"resetCrowdStrikeCustomerId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetDeviceEnrollmentDomain() {
	_jsii_.InvokeVoid(
		d,
		"resetDeviceEnrollmentDomain",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetDiskEncrypted() {
	_jsii_.InvokeVoid(
		d,
		"resetDiskEncrypted",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetKeyTrustLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyTrustLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetManagedDevice() {
	_jsii_.InvokeVoid(
		d,
		"resetManagedDevice",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetOsFirewall() {
	_jsii_.InvokeVoid(
		d,
		"resetOsFirewall",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetOsVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetOsVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetPasswordProtectionWarningTrigger() {
	_jsii_.InvokeVoid(
		d,
		"resetPasswordProtectionWarningTrigger",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetRealtimeUrlCheckMode() {
	_jsii_.InvokeVoid(
		d,
		"resetRealtimeUrlCheckMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetSafeBrowsingProtectionLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetSafeBrowsingProtectionLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetScreenLockSecured() {
	_jsii_.InvokeVoid(
		d,
		"resetScreenLockSecured",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetSiteIsolationEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetSiteIsolationEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetThirdPartyBlockingEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetThirdPartyBlockingEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetWindowsMachineDomain() {
	_jsii_.InvokeVoid(
		d,
		"resetWindowsMachineDomain",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ResetWindowsUserDomain() {
	_jsii_.InvokeVoid(
		d,
		"resetWindowsUserDomain",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

