package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceRelease() *schema.Resource {
	return &schema.Resource{
		Create: resourceReleaseCreate,
		Read:   resourceReleaseRead,
		Delete: resourceReleaseDelete,

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
			"hash": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"jobs": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceReleaseCreate(d *schema.ResourceData, m interface{}) error {
	bdc := m.(*BoshClient)
	name := d.Get("name").(string)
	version := d.Get("version").(string)
	hash := d.Get("hash").(string)
	jobs := d.Get("jobs").([]interface{})

	//cast to []string : rework to be done
	jobslist := make([]string, 0)
	for _, job := range jobs {
		jobslist = append(jobslist, job.(string))
	}

	release := &Release{name, version, hash, jobslist}

	CID, err := bdc.CreateRelease(*release)

	if err != nil {
		return fmt.Errorf("Error : %v", err)
	}
	d.SetId(CID)
	return nil
}

func resourceReleaseRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceReleaseDelete(d *schema.ResourceData, m interface{}) error {
	bdc := m.(*BoshClient)
	CID := d.Id()

	err := bdc.DeleteRelease(CID)

	if err != nil {
		return fmt.Errorf("Error with resource release %s : %v", CID, err)
	}
	return nil
}
