package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeploymentCreate,
		Read:   resourceDeploymentRead,
		Delete: resourceDeploymentDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"releases": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
			"stemcell": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDeploymentCreate(d *schema.ResourceData, m interface{}) error {
	bdc := m.(*BoshClient)
	name := d.Get("name").(string)
	releases := d.Get("releases").([]interface{})
	stemcells := d.Get("stemcell").([]interface{})

	//cast to []string : rework to be done
	releaseslist := make([]string, 0)
	for _, release := range releases {
		releaseslist = append(releaseslist, release.(string))
	}

	stemcellslist := make([]string, 0)
	for _, stemcell := range stemcells {
		stemcellslist = append(stemcellslist, stemcell.(string))
	}

	deployment := &Deployment{name, releaseslist, stemcellslist}

	CID, err := bdc.CreateDeployment(*deployment)

	if err != nil {
		return fmt.Errorf("Error : %v", err)
	}
	d.SetId(CID)
	return nil
}

func resourceDeploymentRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDeploymentDelete(d *schema.ResourceData, m interface{}) error {
	bdc := m.(*BoshClient)
	CID := d.Id()

	err := bdc.DeleteDeployment(CID)

	if err != nil {
		return fmt.Errorf("Error with resource release %s : %v", CID, err)
	}
	return nil
}
