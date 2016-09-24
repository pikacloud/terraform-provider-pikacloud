package pikacloud

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pikacloud/gopikacloud"
)

func TestAccPikacloudZoneRecord_Basic(t *testing.T) {
	var zonerecord gopikacloud.ZoneRecord
	domainName := fmt.Sprintf("foobar-test-terraform-%s.com", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPikacloudZoneRecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckPikacloudZoneRecordBasic, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPikacloudZoneRecordExists("pikacloud_zonerecord.foobar", &zonerecord),
					testAccCheckPikacloudZoneRecordAttributes(&zonerecord, domainName),
					resource.TestCheckResourceAttr(
						"pikacloud_zonerecord.foobar", "name", "www"),
				),
			},
		},
	})
}

func testAccCheckPikacloudZoneRecordDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*gopikacloud.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "pikacloud_record" {
			continue
		}
		zone := rs.Primary.Attributes["zone"]
		_, err := client.ZoneRecord(zone, rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Zone record still exists")
		}
	}

	return nil
}

func testAccCheckPikacloudZoneRecordAttributes(zonerecord *gopikacloud.ZoneRecord, domainName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if zonerecord.Ipv4 != "127.0.0.1" {
			return fmt.Errorf("Bad IPV4: %s", zonerecord.Ipv4)
		}

		return nil
	}
}

func testAccCheckPikacloudZoneRecordExists(n string, zonerecord *gopikacloud.ZoneRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Zone record ID is set")
		}

		client := testAccProvider.Meta().(*gopikacloud.Client)
		zoneID := rs.Primary.Attributes["zone"]
		foundZoneRecord, err := client.ZoneRecord(zoneID, rs.Primary.ID)
		if err != nil {
			return err
		}

		if strconv.Itoa(foundZoneRecord.ID) != rs.Primary.ID {
			return fmt.Errorf("Zone record not found")
		}

		*zonerecord = foundZoneRecord

		return nil
	}
}

const testAccCheckPikacloudZoneRecordBasic = `
resource "pikacloud_zone" "foobar" {
	domain_name       = "%s"
}

resource "pikacloud_zonerecord" "foobar" {
	zone	= "${pikacloud_zone.foobar.id}"
	rtype	= "A"
	name  = "www"
	ipv4  = "127.0.0.1"
}`
