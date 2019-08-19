package vmworkstation

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VMWS_USER", nil),
				Description: "The user name for VMWare Workstation Pro API REST operations.",
				Sensitive:   true,
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VMWS_PASSWORD", nil),
				Description: "The user password for VMWare Workstation Pro API REST operations.",
				Sensitive:   true,
			},
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VMWS_URL", nil),
				Description: "The URL for connect to the API REST",
			},
			"https": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VMWS_HTTPS", true),
				Description: "When this have set to true the 'url' connect to over https",
			},
			"debug": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VMWS_DEBUG", false),
				Description: "Enable debug for find errors",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"vmworkstation_vms":    resourceVMWSVms(),
			"vmworkstation_folder": resourceVMWSFolder(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"vmworkstation_folder": datasourceVMWSFolder(),
		},
		ConfigureFunc: providerConfigure,
	}
	// if provider.Schema.debug == true {
	// 	log.Printf("[VMWS] Fu: Provider Fi: provider.go Ob: %#v\n", provider)
	// }
	return provider
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config, err := NewConfig(d)
	if err != nil {
		return nil, err
	}
	if d.Get("debug") == true {
		log.Printf("[VMWS] Fu: providerConfigure Fi: provider.go Ob: %#v\n", d)
		log.Printf("[VMWS] Fu: providerConfigure Fi: provider.go Ob: %#v\n", config)
	}
	return config.Client()
}
