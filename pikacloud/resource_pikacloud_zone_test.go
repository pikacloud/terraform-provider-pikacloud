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

func TestAccPikacloudZone_Basic(t *testing.T) {
	var zone gopikacloud.Zone
	domainName := fmt.Sprintf("foobar-test-terraform-%s.com", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPikacloudZoneDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckPikacloudZoneBasic, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPikacloudZoneExists("pikacloud_zone.foobar", &zone),
					testAccCheckPikacloudZoneAttributes(&zone, domainName),
					resource.TestCheckResourceAttr(
						"pikacloud_zone.foobar", "domain_name", domainName),
				),
			},
		},
	})
}

func testAccCheckPikacloudZoneDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*gopikacloud.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "pikacloud_zone" {
			continue
		}

		_, err := client.Zone(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Zone still exists")
		}
	}

	return nil
}

func testAccCheckPikacloudZoneAttributes(zone *gopikacloud.Zone, domainName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if zone.DomainName != domainName {
			return fmt.Errorf("Bad domain name: %s", zone.DomainName)
		}

		return nil
	}
}

func testAccCheckPikacloudZoneExists(n string, zone *gopikacloud.Zone) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Zone ID is set")
		}

		client := testAccProvider.Meta().(*gopikacloud.Client)
		foundZone, err := client.Zone(rs.Primary.ID)
		if err != nil {
			return err
		}

		if strconv.Itoa(foundZone.ID) != rs.Primary.ID {
			return fmt.Errorf("Zone not found")
		}

		*zone = foundZone

		return nil
	}
}

const testAccCheckPikacloudZoneBasic = `
resource "pikacloud_zone" "foobar" {
	domain_name       = "%s"
}`
