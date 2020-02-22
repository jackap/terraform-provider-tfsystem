package tfsystem

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a schema.Provider for tfsystem.
func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"tfsystem_file": resourceFile(),
		},
	}
	return p
}
