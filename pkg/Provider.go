package pkg

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/jitesoft/bitwarden-terraform-provider/pkg/schemas"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider {
		Schema: providerSchema(),
		ResourcesMap: map[string]*schema.Resource{
			"secret": schemas.SecretsResource(),
		},
	}
}

func providerSchema() map[string]*schema.Schema  {
	return map[string]*schema.Schema{
		"name": {
			Description: "Vault name (only used for readability)",
			Type: schema.TypeString,
			Required: true,
		},
		"endpoint": {
			Description: "Uri to the BitWarden api (defaults to SaaS BitWarden)",
			Type: schema.TypeString,
			Required: false,
			Default: "https://api.bitwarden.com",
		},
		"identity-endpoint": {
			Description: "Uri to BitWarden identity api (defaults to SaaS BitWarden)",
			Type: schema.TypeString,
			Required: false,
			Default: "https://identity.bitwarden.com/connect/token",
		},
		"secrets-lifetime": {
			Description: "Global lifetime on secrets in terraform state (defaults to -1, which means that they will not be updated)",
			Type: schema.TypeInt,
			Required: false,
			Default: -1,
		},
	}
}
