package main

//@https://www.terraform.io/docs/extend/writing-custom-providers.html
//@https://github.com/terraform-providers

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		//parametres du provider
		Schema: map[string]*schema.Schema{
			"endpoint_flownet": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FLOWNET_ENDPOINT", nil),
				Description: "Endpoint d'un controller flownet de la forme http://ip:port",
			},
			"user_flownet": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FLOWNET_USER", nil),
				Description: "Utilisateur flownet",
			},
			"password_flownet": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FLOWNET_PASSWORD", nil),
				Description: "Password associe a l'utilisateur flownet",
			},
		},
		//ressources tf exposees aux utilisateurs par le provider
		ResourcesMap: map[string]*schema.Resource{
			"flownet_flow": resourceFlow(),
		},
		//Fonction configurant le client (http pour nous) qui va query le remote controller
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint_flownet").(string)
	//TODO Basic Auth at least
	return NewClient(endpoint)
}
