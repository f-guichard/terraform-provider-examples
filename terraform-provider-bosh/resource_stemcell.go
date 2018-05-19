package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceStemcell() *schema.Resource {
	return &schema.Resource{
		Create: resourceStemcellCreate,
		Read:   resourceStemcellRead,
		Delete: resourceStemcellDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"os": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceStemcellCreate(d *schema.ResourceData, m interface{}) error {
	bdc := m.(*BoshClient)
	name := d.Get("name").(string)
	version := d.Get("version").(string)
	os := d.Get("os").(string)

	stemcell := &Stemcell{name, version, os}

	CID, err := bdc.CreateStemcell(*stemcell)

	if err != nil {
		return fmt.Errorf("Error : %v", err)
	}
	d.SetId(CID)
	return nil
}

func resourceStemcellRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceStemcellDelete(d *schema.ResourceData, m interface{}) error {
	bdc := m.(*BoshClient)
	CID := d.Id()

	err := bdc.DeleteStemcell(CID)

	if err != nil {
		return fmt.Errorf("Error with resource release %s : %v", CID, err)
	}
	return nil
}
