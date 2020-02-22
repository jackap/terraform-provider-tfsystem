package tfsystem

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceFileCreate,
		Read:   resourceFileRead,
		Update: resourceFileUpdate,
		Delete: resourceFileDelete,

		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				///        Computed: true,
			},
		},
	}
}

func resourceFileCreate(d *schema.ResourceData, m interface{}) error {
	path := d.Get("path").(string)
	f, err := os.Create(path)
	if err == nil {
		f.Close()
		d.SetId(path)
		return resourceFileRead(d, m)
	}
	return fmt.Errorf("I could not create the resource")
}

func resourceFileRead(d *schema.ResourceData, m interface{}) error {
	path := d.Get("path").(string)
	d.SetId(path)
	d.Set("path", path)
	return nil
}

func resourceFileUpdate(d *schema.ResourceData, m interface{}) error {
	/*
	  d.Partial(true)

	  if d.HasChange("path") {
	    // Try updating the address
	    if err := updatePath(d, m); err != nil {
	      return err
	    }

	    d.SetPartial("path")
	  }

	  d.Partial(false)
	  return resourceFileRead(d, m)
	*/
	oldPath, newPath := d.GetChange("path")
	if !strings.HasPrefix(newPath.(string), "./") {
		return fmt.Errorf("File path should start with ./")
	}
	oldPath, newPath = d.GetChange("path")
	var err = os.Remove(oldPath.(string))
	if err != nil {
		return err
	}
	d.SetId(newPath.(string))
	d.Set("path", newPath)

	/*         if d.HasChange("path") {
	  // Try updating the address
	  if err := updatePath(d, m); err != nil {
	    return err
	  }

	}*/
	return resourceFileCreate(d, m)
}

func resourceFileDelete(d *schema.ResourceData, m interface{}) error {
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	path := d.Get("path").(string)
	_, fileError := os.Stat(path)
	if fileError != nil {
		return nil
	}
	var err = os.Remove(path)
	return err
}

func updatePath(d *schema.ResourceData, m interface{}) error {
	path := d.Get("path").(string)
	if !strings.HasPrefix(path, "./") {
		return fmt.Errorf("File path should start with ./")
	}
	return nil
}
