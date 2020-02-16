package main

import (
        "fmt"
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
                        },
                },
        }
}

func resourceFileCreate(d *schema.ResourceData, m interface{}) error {
        path := d.Get("path").(string)
        d.SetId(path)
        return resourceFileRead(d, m)
}

func resourceFileRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func resourceFileUpdate(d *schema.ResourceData, m interface{}) error {
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
}

func resourceFileDelete(d *schema.ResourceData, m interface{}) error {
        return nil
}

func updatePath(d *schema.ResourceData, m interface{}) error {
         path := d.Get("path").(string)
        if !strings.HasPrefix(path, "./") {
          return fmt.Errorf("File path should start with ./")
        }
         d.SetId(path)
         return nil
}

