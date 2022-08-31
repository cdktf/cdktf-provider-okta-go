//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt okta Provider for Terraform CDK (cdktf)
package okta

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataOktaEmailCustomizationsEmailCustomizationsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataOktaEmailCustomizationsEmailCustomizationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataOktaEmailCustomizationsEmailCustomizationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataOktaEmailCustomizationsEmailCustomizationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataOktaEmailCustomizationsEmailCustomizationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataOktaEmailCustomizationsEmailCustomizationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
