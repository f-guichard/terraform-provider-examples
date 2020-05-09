package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceBasicat() *schema.Resource {
	return &schema.Resource{
		Create: resourceBasicatCreate,
		Read:   resourceBasicatRead,
		Delete: resourceBasicatDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBasicatCreate(d *schema.ResourceData, m interface{}) error {
	cl := m.(*Client)
	name := d.Get("name").(string)

	basicat := &Basicat{name}

	id, err := cl.CreateObjectBasicat(*basicat)

	if err != nil {
		return fmt.Errorf("Error : %v", err)
	}
	d.SetId(id)
	return nil
}

func resourceBasicatRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceBasicatDelete(d *schema.ResourceData, m interface{}) error {
	cl := m.(*Client)
	id := d.Id()

	err := cl.DeleteObject(id)

	if err != nil {
		return fmt.Errorf("Error with resource %s : %v", id, err)
	}
	return nil
}
