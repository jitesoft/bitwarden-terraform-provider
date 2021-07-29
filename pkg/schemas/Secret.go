package schemas

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

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
			"name":  nil,
			"path":  nil,
			"value": nil,
		},
	}
}

func createSecret(d *schema.ResourceData, meta interface{}) error { return nil }
func readSecret(d *schema.ResourceData, meta interface{}) error   { return nil }
func updateSecret(d *schema.ResourceData, meta interface{}) error { return nil }
func deleteSecret(d *schema.ResourceData, meta interface{}) error { return nil }
func existSecret(d *schema.ResourceData, meta interface{}) (bool, error)  { return false, nil }
