// Prebuilt okta Provider for Terraform CDK (cdktf)
package okta

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-okta-go/okta/v2/jsii"

	"github.com/hashicorp/cdktf-provider-okta-go/okta/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicyRuleSignonFactorSequenceSecondaryCriteriaList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) PolicyRuleSignonFactorSequenceSecondaryCriteriaOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PolicyRuleSignonFactorSequenceSecondaryCriteriaList
type jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewPolicyRuleSignonFactorSequenceSecondaryCriteriaList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) PolicyRuleSignonFactorSequenceSecondaryCriteriaList {
	_init_.Initialize()

	j := jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList{}

	_jsii_.Create(
		"@cdktf/provider-okta.PolicyRuleSignonFactorSequenceSecondaryCriteriaList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewPolicyRuleSignonFactorSequenceSecondaryCriteriaList_Override(p PolicyRuleSignonFactorSequenceSecondaryCriteriaList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.PolicyRuleSignonFactorSequenceSecondaryCriteriaList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		p,
	)
}

func (j *jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (p *jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList) Get(index *float64) PolicyRuleSignonFactorSequenceSecondaryCriteriaOutputReference {
	var returns PolicyRuleSignonFactorSequenceSecondaryCriteriaOutputReference

	_jsii_.Invoke(
		p,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
