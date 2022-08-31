//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt okta Provider for Terraform CDK (cdktf)
package okta

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GroupSchemaPropertyArrayOneOfList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GroupSchemaPropertyArrayOneOfList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GroupSchemaPropertyArrayOneOfList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GroupSchemaPropertyArrayOneOfList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GroupSchemaPropertyArrayOneOfList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GroupSchemaPropertyArrayOneOfList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGroupSchemaPropertyArrayOneOfListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
