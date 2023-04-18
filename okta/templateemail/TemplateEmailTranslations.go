package templateemail


type TemplateEmailTranslations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.46.0/docs/resources/template_email#language TemplateEmail#language}.
	Language *string `field:"required" json:"language" yaml:"language"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.46.0/docs/resources/template_email#subject TemplateEmail#subject}.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.46.0/docs/resources/template_email#template TemplateEmail#template}.
	Template *string `field:"required" json:"template" yaml:"template"`
}

