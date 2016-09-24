package pikacloud

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pikacloud/gopikacloud"
)

func resourcePikacloudZoneRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourcePikacloudZoneRecordCreate,
		Read:   resourcePikacloudZoneRecordRead,
		Delete: resourcePikacloudZoneRecordDelete,

		Schema: map[string]*schema.Schema{
			"zone": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"rtype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ipv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  "1800",
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourcePikacloudZoneRecordCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gopikacloud.Client)

	// Create the new zone
	newZoneRecord := &gopikacloud.ZoneRecord{
		Rtype:    d.Get("rtype").(string),
		Name:     d.Get("name").(string),
		Ipv4:     d.Get("ipv4").(string),
		Ipv6:     d.Get("ipv6").(string),
		Hostname: d.Get("hostname").(string),
		Content:  d.Get("content").(string),
		Priority: d.Get("priority").(int),
		TTL:      d.Get("ttl").(int),
	}

	log.Printf("[DEBUG] Pikacloud zone record create configuration: %#v", newZoneRecord)

	zonerecord, err := client.CreateZoneRecord(d.Get("zone").(int), newZoneRecord)

	if err != nil {
		return fmt.Errorf("Failed to create Pikacloud zone record: %s", err)
	}
	d.SetId(strconv.Itoa(zonerecord.ID))
	log.Printf("[INFO] zone record ID: %s", d.Id())

	return resourcePikacloudZoneRecordRead(d, meta)
}

func resourcePikacloudZoneRecordRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gopikacloud.Client)

	zonerecord, err := client.ZoneRecord(d.Get("zone").(int), d.Id())
	if err != nil {
		return fmt.Errorf("Couldn't find Pikacloud zone record: %s", err)
	}

	d.Set("id", zonerecord.ID)
	d.Set("zone", zonerecord.ZoneID)
	d.Set("rtype", zonerecord.Rtype)
	d.Set("name", zonerecord.Name)
	d.Set("ipv4", zonerecord.Ipv4)
	d.Set("ipv6", zonerecord.Ipv6)
	d.Set("hostname", zonerecord.Hostname)
	d.Set("content", zonerecord.Content)
	d.Set("priority", zonerecord.Priority)
	d.Set("ttl", zonerecord.TTL)

	return nil
}

func resourcePikacloudZoneRecordDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gopikacloud.Client)
	zonerecord := gopikacloud.ZoneRecord{ZoneID: d.Get("zone").(int), ID: d.Get("id").(int)}

	log.Printf("[INFO] Deleting Pikacloud zone record: %s", d.Id())

	err := zonerecord.Delete(client)

	if err != nil {
		return fmt.Errorf("Error deleting Pikacloud zone record: %s", err)
	}

	return nil
}
