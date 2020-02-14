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
			"endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ENDPOINT", nil),
				Description: "Endpoint d'un controller netstat de la forme http://ip:port",
			},
			"user": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("USER", nil),
				Description: "Utilisateur netstat",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PASSWORD", nil),
				Description: "Password associe a l'utilisateur netstat",
			},
		},
		//ressources tf exposees aux utilisateurs par le provider
		ResourcesMap: map[string]*schema.Resource{
			"dns4hbx_zone":   resourceZone(),
			"dns4hbx_record": resourceRecord(),
		},
		//Fonction configurant le client (http pour nous) qui va query le remote controller
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint").(string)
	//TODO Basic Auth at least
	return NewClient(endpoint)
}
