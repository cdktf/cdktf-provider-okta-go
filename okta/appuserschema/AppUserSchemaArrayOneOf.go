package appuserschema


type AppUserSchemaArrayOneOf struct {
	// Enum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.46.0/docs/resources/app_user_schema#const AppUserSchema#const}
	Const *string `field:"required" json:"const" yaml:"const"`
	// Enum title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.46.0/docs/resources/app_user_schema#title AppUserSchema#title}
	Title *string `field:"required" json:"title" yaml:"title"`
}

