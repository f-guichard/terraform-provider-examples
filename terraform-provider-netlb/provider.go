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
			"endpoint_networklb": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETWORKLB_ENDPOINT", nil),
				Description: "Endpoint d'un controller network-lb de la forme http://ip:port",
			},
			"user_networklb": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETWORKLB_USER", nil),
				Description: "Utilisateur network-lb",
			},
			"password_networklb": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETWORKLB_PASSWORD", nil),
				Description: "Password associe a l'utilisateur network-lb",
			},
		},
		//ressources tf exposees aux utilisateurs par le provider
		ResourcesMap: map[string]*schema.Resource{
			"networklb_vip": resourceVip(),
		},
		//Fonction configurant le client (http pour nous) qui va query le remote controller
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint_networklb").(string)
	//TODO Basic Auth at least
	return NewClient(endpoint)
}
