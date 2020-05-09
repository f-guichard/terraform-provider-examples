package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationCreate,
		Read:   resourceApplicationRead,
		Delete: resourceApplicationDelete,

		Schema: map[string]*schema.Schema{
			"modapp":  &schema.Schema{
                Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApplicationCreate(d *schema.ResourceData, m interface{}) error {
	cl := m.(*Client)
	modapp := d.Get("modapp").(string)
	name := d.Get("name").(string)

	app := &Application{modapp, name}

	id, err := cl.CreateObjectApplication(*app)

	if err != nil {
		return fmt.Errorf("Error : %v", err)
	}
	d.SetId(id)
	return nil
}

func resourceApplicationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceApplicationDelete(d *schema.ResourceData, m interface{}) error {
	cl := m.(*Client)
	id := d.Id()

	err := cl.DeleteObject(id)

	if err != nil {
		return fmt.Errorf("Error with resource %s : %v", id, err)
	}
	return nil
}
