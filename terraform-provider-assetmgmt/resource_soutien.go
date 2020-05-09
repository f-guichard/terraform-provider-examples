package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSoutien() *schema.Resource {
	return &schema.Resource{
		Create: resourceSoutienCreate,
		Read:   resourceSoutienRead,
		Delete: resourceSoutienDelete,

		Schema: map[string]*schema.Schema{
			"exploit": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"support": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resp_exploit": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resp_support": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"etat": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"eds_peit": &schema.Schema{
                Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"eds_sat": &schema.Schema{
                Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSoutienCreate(d *schema.ResourceData, m interface{}) error {
	cl := m.(*Client)
	exploit := d.Get("exploit").(string)
	support := d.Get("support").(string)
	rexploit := d.Get("resp_exploit").(string)
	rsupport := d.Get("resp_support").(string)
	etat := d.Get("etat").(string)
	edspeit := d.Get("eds_peit").(string)
	edssat := d.Get("eds_sat").(string)

	soutien := &Soutien{exploit, support, rexploit, rsupport, etat, edspeit, edssat}

	id, err := cl.CreateObjectSoutien(*soutien)

	if err != nil {
		return fmt.Errorf("Error : %v", err)
	}
	d.SetId(id)
	return nil
}

func resourceSoutienRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSoutienDelete(d *schema.ResourceData, m interface{}) error {
	cl := m.(*Client)
	id := d.Id()

	err := cl.DeleteObject(id)

	if err != nil {
		return fmt.Errorf("Error with resource %s : %v", id, err)
	}
	return nil
}
