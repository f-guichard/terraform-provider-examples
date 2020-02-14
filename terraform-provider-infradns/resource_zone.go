package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceZoneCreate,
		Read:   resourceZoneRead,
		Delete: resourceZoneDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceZoneCreate(d *schema.ResourceData, m interface{}) error {
	cl := m.(*Client)
	name := d.Get("name").(string)

	zone := &Zone{name}

	id, err := cl.CreateObjectZone(*zone)

	if err != nil {
		return fmt.Errorf("Error : %v", err)
	}
	d.SetId(id)
	return nil
}

func resourceZoneRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceZoneDelete(d *schema.ResourceData, m interface{}) error {
	cl := m.(*Client)
	id := d.Id()

	err := cl.DeleteObject(id)

	if err != nil {
		return fmt.Errorf("Error with resource %s : %v", id, err)
	}
	return nil
}
