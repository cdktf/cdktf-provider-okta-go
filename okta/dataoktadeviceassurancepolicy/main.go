// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataoktadeviceassurancepolicy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicy",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "diskEncryptionType", GoGetter: "DiskEncryptionType"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "jailbreak", GoGetter: "Jailbreak"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "osVersion", GoGetter: "OsVersion"},
			_jsii_.MemberProperty{JsiiProperty: "osVersionConstraint", GoGetter: "OsVersionConstraint"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "platform", GoGetter: "Platform"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putThirdPartySignalProvider", GoMethod: "PutThirdPartySignalProvider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecureHardwarePresent", GoMethod: "ResetSecureHardwarePresent"},
			_jsii_.MemberMethod{JsiiMethod: "resetThirdPartySignalProvider", GoMethod: "ResetThirdPartySignalProvider"},
			_jsii_.MemberProperty{JsiiProperty: "screenlockType", GoGetter: "ScreenlockType"},
			_jsii_.MemberProperty{JsiiProperty: "secureHardwarePresent", GoGetter: "SecureHardwarePresent"},
			_jsii_.MemberProperty{JsiiProperty: "secureHardwarePresentInput", GoGetter: "SecureHardwarePresentInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "thirdPartySignalProvider", GoGetter: "ThirdPartySignalProvider"},
			_jsii_.MemberProperty{JsiiProperty: "thirdPartySignalProviderInput", GoGetter: "ThirdPartySignalProviderInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaDeviceAssurancePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyConfig",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyDiskEncryptionType",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyDiskEncryptionType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyDiskEncryptionTypeOutputReference",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyDiskEncryptionTypeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaDeviceAssurancePolicyDiskEncryptionTypeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyOsVersion",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyOsVersion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyOsVersionConstraint",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyOsVersionConstraint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyOsVersionConstraintDynamicVersionRequirement",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyOsVersionConstraintDynamicVersionRequirement)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyOsVersionConstraintDynamicVersionRequirementOutputReference",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyOsVersionConstraintDynamicVersionRequirementOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "distanceFromLatestMajor", GoGetter: "DistanceFromLatestMajor"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "latestSecurityPatch", GoGetter: "LatestSecurityPatch"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaDeviceAssurancePolicyOsVersionConstraintDynamicVersionRequirementOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyOsVersionConstraintList",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyOsVersionConstraintList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaDeviceAssurancePolicyOsVersionConstraintList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyOsVersionConstraintOutputReference",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyOsVersionConstraintOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicVersionRequirement", GoGetter: "DynamicVersionRequirement"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "majorVersionConstraint", GoGetter: "MajorVersionConstraint"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaDeviceAssurancePolicyOsVersionConstraintOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyOsVersionDynamicVersionRequirement",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyOsVersionDynamicVersionRequirement)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyOsVersionDynamicVersionRequirementOutputReference",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyOsVersionDynamicVersionRequirementOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "distanceFromLatestMajor", GoGetter: "DistanceFromLatestMajor"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "latestSecurityPatch", GoGetter: "LatestSecurityPatch"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaDeviceAssurancePolicyOsVersionDynamicVersionRequirementOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyOsVersionOutputReference",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyOsVersionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicVersionRequirement", GoGetter: "DynamicVersionRequirement"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "minimum", GoGetter: "Minimum"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaDeviceAssurancePolicyOsVersionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyScreenlockType",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyScreenlockType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyScreenlockTypeOutputReference",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyScreenlockTypeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaDeviceAssurancePolicyScreenlockTypeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyThirdPartySignalProvider",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyThirdPartySignalProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtc",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcBrowserVersion",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcBrowserVersion)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcBrowserVersionOutputReference",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcBrowserVersionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "minimum", GoGetter: "Minimum"},
			_jsii_.MemberProperty{JsiiProperty: "minimumInput", GoGetter: "MinimumInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinimum", GoMethod: "ResetMinimum"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcBrowserVersionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOsVersion",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOsVersion)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOsVersionOutputReference",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOsVersionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "minimum", GoGetter: "Minimum"},
			_jsii_.MemberProperty{JsiiProperty: "minimumInput", GoGetter: "MinimumInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinimum", GoMethod: "ResetMinimum"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOsVersionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowScreenLock", GoGetter: "AllowScreenLock"},
			_jsii_.MemberProperty{JsiiProperty: "allowScreenLockInput", GoGetter: "AllowScreenLockInput"},
			_jsii_.MemberProperty{JsiiProperty: "browserVersion", GoGetter: "BrowserVersion"},
			_jsii_.MemberProperty{JsiiProperty: "browserVersionInput", GoGetter: "BrowserVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "builtInDnsClientEnabled", GoGetter: "BuiltInDnsClientEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "builtInDnsClientEnabledInput", GoGetter: "BuiltInDnsClientEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "chromeRemoteDesktopAppBlocked", GoGetter: "ChromeRemoteDesktopAppBlocked"},
			_jsii_.MemberProperty{JsiiProperty: "chromeRemoteDesktopAppBlockedInput", GoGetter: "ChromeRemoteDesktopAppBlockedInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "crowdStrikeAgentId", GoGetter: "CrowdStrikeAgentId"},
			_jsii_.MemberProperty{JsiiProperty: "crowdStrikeAgentIdInput", GoGetter: "CrowdStrikeAgentIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "crowdStrikeCustomerId", GoGetter: "CrowdStrikeCustomerId"},
			_jsii_.MemberProperty{JsiiProperty: "crowdStrikeCustomerIdInput", GoGetter: "CrowdStrikeCustomerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceEnrollmentDomain", GoGetter: "DeviceEnrollmentDomain"},
			_jsii_.MemberProperty{JsiiProperty: "deviceEnrollmentDomainInput", GoGetter: "DeviceEnrollmentDomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "diskEncrypted", GoGetter: "DiskEncrypted"},
			_jsii_.MemberProperty{JsiiProperty: "diskEncryptedInput", GoGetter: "DiskEncryptedInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyTrustLevel", GoGetter: "KeyTrustLevel"},
			_jsii_.MemberProperty{JsiiProperty: "keyTrustLevelInput", GoGetter: "KeyTrustLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "managedDevice", GoGetter: "ManagedDevice"},
			_jsii_.MemberProperty{JsiiProperty: "managedDeviceInput", GoGetter: "ManagedDeviceInput"},
			_jsii_.MemberProperty{JsiiProperty: "osFirewall", GoGetter: "OsFirewall"},
			_jsii_.MemberProperty{JsiiProperty: "osFirewallInput", GoGetter: "OsFirewallInput"},
			_jsii_.MemberProperty{JsiiProperty: "osVersion", GoGetter: "OsVersion"},
			_jsii_.MemberProperty{JsiiProperty: "osVersionInput", GoGetter: "OsVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordProtectionWarningTrigger", GoGetter: "PasswordProtectionWarningTrigger"},
			_jsii_.MemberProperty{JsiiProperty: "passwordProtectionWarningTriggerInput", GoGetter: "PasswordProtectionWarningTriggerInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBrowserVersion", GoMethod: "PutBrowserVersion"},
			_jsii_.MemberMethod{JsiiMethod: "putOsVersion", GoMethod: "PutOsVersion"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeUrlCheckMode", GoGetter: "RealtimeUrlCheckMode"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeUrlCheckModeInput", GoGetter: "RealtimeUrlCheckModeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowScreenLock", GoMethod: "ResetAllowScreenLock"},
			_jsii_.MemberMethod{JsiiMethod: "resetBrowserVersion", GoMethod: "ResetBrowserVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetBuiltInDnsClientEnabled", GoMethod: "ResetBuiltInDnsClientEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetChromeRemoteDesktopAppBlocked", GoMethod: "ResetChromeRemoteDesktopAppBlocked"},
			_jsii_.MemberMethod{JsiiMethod: "resetCrowdStrikeAgentId", GoMethod: "ResetCrowdStrikeAgentId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCrowdStrikeCustomerId", GoMethod: "ResetCrowdStrikeCustomerId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceEnrollmentDomain", GoMethod: "ResetDeviceEnrollmentDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskEncrypted", GoMethod: "ResetDiskEncrypted"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyTrustLevel", GoMethod: "ResetKeyTrustLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetManagedDevice", GoMethod: "ResetManagedDevice"},
			_jsii_.MemberMethod{JsiiMethod: "resetOsFirewall", GoMethod: "ResetOsFirewall"},
			_jsii_.MemberMethod{JsiiMethod: "resetOsVersion", GoMethod: "ResetOsVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordProtectionWarningTrigger", GoMethod: "ResetPasswordProtectionWarningTrigger"},
			_jsii_.MemberMethod{JsiiMethod: "resetRealtimeUrlCheckMode", GoMethod: "ResetRealtimeUrlCheckMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetSafeBrowsingProtectionLevel", GoMethod: "ResetSafeBrowsingProtectionLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetScreenLockSecured", GoMethod: "ResetScreenLockSecured"},
			_jsii_.MemberMethod{JsiiMethod: "resetSiteIsolationEnabled", GoMethod: "ResetSiteIsolationEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetThirdPartyBlockingEnabled", GoMethod: "ResetThirdPartyBlockingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetWindowsMachineDomain", GoMethod: "ResetWindowsMachineDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetWindowsUserDomain", GoMethod: "ResetWindowsUserDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "safeBrowsingProtectionLevel", GoGetter: "SafeBrowsingProtectionLevel"},
			_jsii_.MemberProperty{JsiiProperty: "safeBrowsingProtectionLevelInput", GoGetter: "SafeBrowsingProtectionLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "screenLockSecured", GoGetter: "ScreenLockSecured"},
			_jsii_.MemberProperty{JsiiProperty: "screenLockSecuredInput", GoGetter: "ScreenLockSecuredInput"},
			_jsii_.MemberProperty{JsiiProperty: "siteIsolationEnabled", GoGetter: "SiteIsolationEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "siteIsolationEnabledInput", GoGetter: "SiteIsolationEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "thirdPartyBlockingEnabled", GoGetter: "ThirdPartyBlockingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "thirdPartyBlockingEnabledInput", GoGetter: "ThirdPartyBlockingEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "windowsMachineDomain", GoGetter: "WindowsMachineDomain"},
			_jsii_.MemberProperty{JsiiProperty: "windowsMachineDomainInput", GoGetter: "WindowsMachineDomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "windowsUserDomain", GoGetter: "WindowsUserDomain"},
			_jsii_.MemberProperty{JsiiProperty: "windowsUserDomainInput", GoGetter: "WindowsUserDomainInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderDtcOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaDeviceAssurancePolicy.DataOktaDeviceAssurancePolicyThirdPartySignalProviderOutputReference",
		reflect.TypeOf((*DataOktaDeviceAssurancePolicyThirdPartySignalProviderOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dtc", GoGetter: "Dtc"},
			_jsii_.MemberProperty{JsiiProperty: "dtcInput", GoGetter: "DtcInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putDtc", GoMethod: "PutDtc"},
			_jsii_.MemberMethod{JsiiMethod: "resetDtc", GoMethod: "ResetDtc"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaDeviceAssurancePolicyThirdPartySignalProviderOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
