package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceRecordCreate,
		Read:   resourceRecordRead,
		Delete: resourceRecordDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRecordCreate(d *schema.ResourceData, m interface{}) error {
	cl := m.(*Client)
	name := d.Get("name").(string)
	rtype := d.Get("type").(string)
	address := d.Get("address").(string)
	zone := d.Get("zone").(string)

	record := &Record{name, rtype, address, zone}

	id, err := cl.CreateObjectRecord(*record)

	if err != nil {
		return fmt.Errorf("Error : %v", err)
	}
	d.SetId(id)
	return nil
}

func resourceRecordRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceRecordDelete(d *schema.ResourceData, m interface{}) error {
	cl := m.(*Client)
	id := d.Id()

	err := cl.DeleteObject(id)

	if err != nil {
		return fmt.Errorf("Error with resource %s : %v", id, err)
	}
	return nil
}
