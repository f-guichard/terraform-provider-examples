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
			"endpoint_boshdirector": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BOSH_ENDPOINT", nil),
				Description: "Endpoint d'un director bosh de la forme http://ip:port",
			},
			"user_boshdirector": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BOSH_USER", nil),
				Description: "Utilisateur bosh",
			},
			"password_boshdirector": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BOSH_PASSWORD", nil),
				Description: "Password associe a l'utilisateur bosh",
			},
		},
		//ressources tf exposees aux utilisateurs par le provider
		ResourcesMap: map[string]*schema.Resource{
			"bosh_stemcell":   resourceStemcell(),
			"bosh_release":    resourceRelease(),
			"bosh_deployment": resourceDeployment(),
		},
		//Fonction configurant le client (http pour nous) qui va query le remote controller
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint_boshdirector").(string)
	//TODO Basic Auth at least
	return NewClient(endpoint)
}
