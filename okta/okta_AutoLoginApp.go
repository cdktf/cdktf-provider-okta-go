// Prebuilt okta Provider for Terraform CDK (cdktf)
package okta

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-okta-go/okta/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-okta-go/okta/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/okta/r/auto_login_app okta_auto_login_app}.
type AutoLoginApp interface {
	cdktf.TerraformResource
	AccessibilityErrorRedirectUrl() *string
	SetAccessibilityErrorRedirectUrl(val *string)
	AccessibilityErrorRedirectUrlInput() *string
	AccessibilityLoginRedirectUrl() *string
	SetAccessibilityLoginRedirectUrl(val *string)
	AccessibilityLoginRedirectUrlInput() *string
	AccessibilitySelfService() interface{}
	SetAccessibilitySelfService(val interface{})
	AccessibilitySelfServiceInput() interface{}
	AdminNote() *string
	SetAdminNote(val *string)
	AdminNoteInput() *string
	AppLinksJson() *string
	SetAppLinksJson(val *string)
	AppLinksJsonInput() *string
	AppSettingsJson() *string
	SetAppSettingsJson(val *string)
	AppSettingsJsonInput() *string
	AutoSubmitToolbar() interface{}
	SetAutoSubmitToolbar(val interface{})
	AutoSubmitToolbarInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CredentialsScheme() *string
	SetCredentialsScheme(val *string)
	CredentialsSchemeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnduserNote() *string
	SetEnduserNote(val *string)
	EnduserNoteInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Groups() *[]*string
	SetGroups(val *[]*string)
	GroupsInput() *[]*string
	HideIos() interface{}
	SetHideIos(val interface{})
	HideIosInput() interface{}
	HideWeb() interface{}
	SetHideWeb(val interface{})
	HideWebInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Logo() *string
	SetLogo(val *string)
	LogoInput() *string
	LogoUrl() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	PreconfiguredApp() *string
	SetPreconfiguredApp(val *string)
	PreconfiguredAppInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RevealPassword() interface{}
	SetRevealPassword(val interface{})
	RevealPasswordInput() interface{}
	SharedPassword() *string
	SetSharedPassword(val *string)
	SharedPasswordInput() *string
	SharedUsername() *string
	SetSharedUsername(val *string)
	SharedUsernameInput() *string
	SignOnMode() *string
	SignOnRedirectUrl() *string
	SetSignOnRedirectUrl(val *string)
	SignOnRedirectUrlInput() *string
	SignOnUrl() *string
	SetSignOnUrl(val *string)
	SignOnUrlInput() *string
	SkipGroups() interface{}
	SetSkipGroups(val interface{})
	SkipGroupsInput() interface{}
	SkipUsers() interface{}
	SetSkipUsers(val interface{})
	SkipUsersInput() interface{}
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserNameTemplate() *string
	SetUserNameTemplate(val *string)
	UserNameTemplateInput() *string
	UserNameTemplatePushStatus() *string
	SetUserNameTemplatePushStatus(val *string)
	UserNameTemplatePushStatusInput() *string
	UserNameTemplateSuffix() *string
	SetUserNameTemplateSuffix(val *string)
	UserNameTemplateSuffixInput() *string
	UserNameTemplateType() *string
	SetUserNameTemplateType(val *string)
	UserNameTemplateTypeInput() *string
	Users() AutoLoginAppUsersList
	UsersInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutUsers(value interface{})
	ResetAccessibilityErrorRedirectUrl()
	ResetAccessibilityLoginRedirectUrl()
	ResetAccessibilitySelfService()
	ResetAdminNote()
	ResetAppLinksJson()
	ResetAppSettingsJson()
	ResetAutoSubmitToolbar()
	ResetCredentialsScheme()
	ResetEnduserNote()
	ResetGroups()
	ResetHideIos()
	ResetHideWeb()
	ResetId()
	ResetLogo()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreconfiguredApp()
	ResetRevealPassword()
	ResetSharedPassword()
	ResetSharedUsername()
	ResetSignOnRedirectUrl()
	ResetSignOnUrl()
	ResetSkipGroups()
	ResetSkipUsers()
	ResetStatus()
	ResetUserNameTemplate()
	ResetUserNameTemplatePushStatus()
	ResetUserNameTemplateSuffix()
	ResetUserNameTemplateType()
	ResetUsers()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AutoLoginApp
type jsiiProxy_AutoLoginApp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AutoLoginApp) AccessibilityErrorRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) AccessibilityErrorRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) AccessibilityLoginRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) AccessibilityLoginRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) AccessibilitySelfService() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) AccessibilitySelfServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) AdminNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) AdminNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) AppLinksJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) AppLinksJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) AppSettingsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSettingsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) AppSettingsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSettingsJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) AutoSubmitToolbar() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) AutoSubmitToolbarInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) CredentialsScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) CredentialsSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) EnduserNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) EnduserNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) GroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) HideIos() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) HideIosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) HideWeb() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWeb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) HideWebInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWebInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) Logo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) LogoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) LogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) PreconfiguredApp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconfiguredApp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) PreconfiguredAppInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconfiguredAppInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) RevealPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revealPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) RevealPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revealPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) SharedPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) SharedPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) SharedUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) SharedUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) SignOnMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) SignOnRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) SignOnRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) SignOnUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) SignOnUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) SkipGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) SkipGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) SkipUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) SkipUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) UserNameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) UserNameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) UserNameTemplatePushStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) UserNameTemplatePushStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) UserNameTemplateSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) UserNameTemplateSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) UserNameTemplateType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) UserNameTemplateTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) Users() AutoLoginAppUsersList {
	var returns AutoLoginAppUsersList
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoLoginApp) UsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/okta/r/auto_login_app okta_auto_login_app} Resource.
func NewAutoLoginApp(scope constructs.Construct, id *string, config *AutoLoginAppConfig) AutoLoginApp {
	_init_.Initialize()

	j := jsiiProxy_AutoLoginApp{}

	_jsii_.Create(
		"@cdktf/provider-okta.AutoLoginApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/okta/r/auto_login_app okta_auto_login_app} Resource.
func NewAutoLoginApp_Override(a AutoLoginApp, scope constructs.Construct, id *string, config *AutoLoginAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.AutoLoginApp",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetAccessibilityErrorRedirectUrl(val *string) {
	_jsii_.Set(
		j,
		"accessibilityErrorRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetAccessibilityLoginRedirectUrl(val *string) {
	_jsii_.Set(
		j,
		"accessibilityLoginRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetAccessibilitySelfService(val interface{}) {
	_jsii_.Set(
		j,
		"accessibilitySelfService",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetAdminNote(val *string) {
	_jsii_.Set(
		j,
		"adminNote",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetAppLinksJson(val *string) {
	_jsii_.Set(
		j,
		"appLinksJson",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetAppSettingsJson(val *string) {
	_jsii_.Set(
		j,
		"appSettingsJson",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetAutoSubmitToolbar(val interface{}) {
	_jsii_.Set(
		j,
		"autoSubmitToolbar",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetCredentialsScheme(val *string) {
	_jsii_.Set(
		j,
		"credentialsScheme",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetEnduserNote(val *string) {
	_jsii_.Set(
		j,
		"enduserNote",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"groups",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetHideIos(val interface{}) {
	_jsii_.Set(
		j,
		"hideIos",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetHideWeb(val interface{}) {
	_jsii_.Set(
		j,
		"hideWeb",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetLabel(val *string) {
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetLogo(val *string) {
	_jsii_.Set(
		j,
		"logo",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetPreconfiguredApp(val *string) {
	_jsii_.Set(
		j,
		"preconfiguredApp",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetRevealPassword(val interface{}) {
	_jsii_.Set(
		j,
		"revealPassword",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetSharedPassword(val *string) {
	_jsii_.Set(
		j,
		"sharedPassword",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetSharedUsername(val *string) {
	_jsii_.Set(
		j,
		"sharedUsername",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetSignOnRedirectUrl(val *string) {
	_jsii_.Set(
		j,
		"signOnRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetSignOnUrl(val *string) {
	_jsii_.Set(
		j,
		"signOnUrl",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetSkipGroups(val interface{}) {
	_jsii_.Set(
		j,
		"skipGroups",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetSkipUsers(val interface{}) {
	_jsii_.Set(
		j,
		"skipUsers",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetUserNameTemplate(val *string) {
	_jsii_.Set(
		j,
		"userNameTemplate",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetUserNameTemplatePushStatus(val *string) {
	_jsii_.Set(
		j,
		"userNameTemplatePushStatus",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetUserNameTemplateSuffix(val *string) {
	_jsii_.Set(
		j,
		"userNameTemplateSuffix",
		val,
	)
}

func (j *jsiiProxy_AutoLoginApp) SetUserNameTemplateType(val *string) {
	_jsii_.Set(
		j,
		"userNameTemplateType",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AutoLoginApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.AutoLoginApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AutoLoginApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.AutoLoginApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AutoLoginApp) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AutoLoginApp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoLoginApp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoLoginApp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoLoginApp) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoLoginApp) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoLoginApp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoLoginApp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoLoginApp) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoLoginApp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoLoginApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoLoginApp) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AutoLoginApp) PutUsers(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putUsers",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetAccessibilityErrorRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilityErrorRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetAccessibilityLoginRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilityLoginRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetAccessibilitySelfService() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilitySelfService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetAdminNote() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminNote",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetAppLinksJson() {
	_jsii_.InvokeVoid(
		a,
		"resetAppLinksJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetAppSettingsJson() {
	_jsii_.InvokeVoid(
		a,
		"resetAppSettingsJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetAutoSubmitToolbar() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoSubmitToolbar",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetCredentialsScheme() {
	_jsii_.InvokeVoid(
		a,
		"resetCredentialsScheme",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetEnduserNote() {
	_jsii_.InvokeVoid(
		a,
		"resetEnduserNote",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetHideIos() {
	_jsii_.InvokeVoid(
		a,
		"resetHideIos",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetHideWeb() {
	_jsii_.InvokeVoid(
		a,
		"resetHideWeb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetLogo() {
	_jsii_.InvokeVoid(
		a,
		"resetLogo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetPreconfiguredApp() {
	_jsii_.InvokeVoid(
		a,
		"resetPreconfiguredApp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetRevealPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetRevealPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetSharedPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetSharedUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetSignOnRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetSignOnRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetSignOnUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetSignOnUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetSkipGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetSkipUsers() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipUsers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetUserNameTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetUserNameTemplatePushStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplatePushStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetUserNameTemplateSuffix() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplateSuffix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetUserNameTemplateType() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplateType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) ResetUsers() {
	_jsii_.InvokeVoid(
		a,
		"resetUsers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoLoginApp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoLoginApp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoLoginApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoLoginApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

