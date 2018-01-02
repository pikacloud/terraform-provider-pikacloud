package pikacloud

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pikacloud/gopikacloud"
)

func resourcePikacloudZone() *schema.Resource {
	return &schema.Resource{
		Create: resourcePikacloudZoneCreate,
		Read:   resourcePikacloudZoneRead,
		Delete: resourcePikacloudZoneDelete,

		Schema: map[string]*schema.Schema{
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourcePikacloudZoneCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gopikacloud.Client)

	// Create the new zone
	newZone := &gopikacloud.Zone{
		DomainName: d.Get("domain_name").(string),
	}

	log.Printf("[DEBUG] Pikacloud zone create configuration: %#v", newZone)

	zone, err := client.CreateZone(newZone)

	if err != nil {
		return fmt.Errorf("Failed to create Pikacloud zone: %s", err)
	}
	d.SetId(strconv.Itoa(zone.ID))
	log.Printf("[INFO] zone ID: %s", d.Id())

	return resourcePikacloudZoneRead(d, meta)
}

func resourcePikacloudZoneRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gopikacloud.Client)
	zoneID, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	zone, err := client.Zone(zoneID)
	if err != nil {
		return fmt.Errorf("Couldn't find Pikacloud zone: %s", err)
	}
	d.Set("domain_name", zone.DomainName)
	d.Set("serial", zone.Serial)
	return nil
}

func resourcePikacloudZoneDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gopikacloud.Client)
	zoneID, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	zone := gopikacloud.Zone{ID: zoneID}
	log.Printf("[INFO] Deleting Pikacloud zone: %s, %s", d.Get("domain_name").(string), d.Id())

	errDelete := zone.Delete(client)

	if errDelete != nil {
		return fmt.Errorf("Error deleting Pikacloud zone: %s", errDelete)
	}

	return nil
}
