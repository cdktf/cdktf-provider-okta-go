//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt okta Provider for Terraform CDK (cdktf)
package okta

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DomainDnsRecordsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DomainDnsRecordsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DomainDnsRecordsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DomainDnsRecordsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DomainDnsRecordsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDomainDnsRecordsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
