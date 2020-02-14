package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVip() *schema.Resource {
	return &schema.Resource{
		Create: resourceVipCreate,
		Read:   resourceVipRead,
		Delete: resourceVipDelete,

		Schema: map[string]*schema.Schema{
			"ipadrress": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"tls_activated": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceVipCreate(d *schema.ResourceData, m interface{}) error {
	nlbc := m.(*NetworklbClient)
	ipadrress := d.Get("ipadrress").(string)
	port := d.Get("port").(int)
	tlsActivated := d.Get("tls_activated").(bool)

	vip := &Vip{ipadrress, port, tlsActivated}

	vipID, err := nlbc.CreateVip(*vip)

	if err != nil {
		return fmt.Errorf("Error : %v", err)
	}
	d.SetId(vipID)
	return nil
}

func resourceVipRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceVipDelete(d *schema.ResourceData, m interface{}) error {
	nlbc := m.(*NetworklbClient)
	vipID := d.Id()

	err := nlbc.DeleteVip(vipID)

	if err != nil {
		return fmt.Errorf("Error with resource %s : %v", vipID, err)
	}
	return nil
}
