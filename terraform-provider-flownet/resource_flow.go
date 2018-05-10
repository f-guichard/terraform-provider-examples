package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFlow() *schema.Resource {
	return &schema.Resource{
		Create: resourceFlowCreate,
		Read:   resourceFlowRead,
		Update: resourceFlowUpdate,
		Delete: resourceFlowDelete,

		Schema: map[string]*schema.Schema{
			"cidr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourceFlowCreate(d *schema.ResourceData, m interface{}) error {
	fnc := m.(*FlownetClient)
	cidr := d.Get("cidr").(string)
	port := d.Get("port").(string)
	description := d.Get("description").(string)

	flow := &Flow{cidr, port, description}

	flowID, err := fnc.CreateFlow(*flow)

	if err != nil {
		return fmt.Errorf("Error : %v", err)
	}
	d.SetId(flowID)
	return nil
}

func resourceFlowRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFlowUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFlowDelete(d *schema.ResourceData, m interface{}) error {
	fnc := m.(*FlownetClient)
	flowID := d.Id()

	err := fnc.DeleteFlow(flowID)

	if err != nil {
		return fmt.Errorf("Error with resource %s : %v", flowID, err)
	}
	return nil
}
