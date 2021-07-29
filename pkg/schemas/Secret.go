package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/jitesoft/bitwarden-terraform-provider/pkg/api"
)

func SecretsResource() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSecret,
		Read:          readSecret,
		Update:        updateSecret,
		Delete:        deleteSecret,
		Exists:        existSecret,
		Description:   "Resource for a single BitWarden item/entry",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the item",
				Sensitive:   true,
			},
			"path":  {
				Type: schema.TypeString,
				Required: true,
				Sensitive: true,
				Description: "Path to the entry in the vault",
			},
			"value": {
				Type: schema.TypeString,
				Required: true,
				Sensitive: true,
				Description: "Value of the item",
			},
			"lifetime": {
				Type: schema.TypeInt,
				Required: false,
				Sensitive: false,
				Description: "Lifetime in seconds of the given entry (defaults to provider globally set lifetime)",
			},
		},
	}
}

func createSecret(d *schema.ResourceData, meta interface{}) error {
	err := api.CreateEntry()
	if err != nil {
		return err
	}

	// Save.
	return err
}

func readSecret(d *schema.ResourceData, meta interface{}) error   {
	obj, err := api.GetEntry()
	if err != nil {
		return err
	}

	err = d.Set("value", obj.Value)
	if err != nil {
		return err
	}

	// Save
	return nil
}

func updateSecret(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteSecret(d *schema.ResourceData, meta interface{}) error { return nil }

func existSecret(d *schema.ResourceData, meta interface{}) (bool, error)  { return false, nil }
