package tfsystem

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

/*
func init() {
	resource.AddTestSweepers("digitalocean_floating_ip", &resource.Sweeper{
		Name: "digitalocean_floating_ip",
		F:    testSweepFloatingIps,
	})

}

func testSweepFloatingIps(region string) error {
	meta, err := sharedConfigForRegion(region)
	if err != nil {
		return err
	}

	client := meta.(*CombinedConfig).godoClient()

	ips, _, err := client.FloatingIPs.List(context.Background(), nil)
	if err != nil {
		return err
	}

	for _, ip := range ips {
		if _, err := client.FloatingIPs.Delete(context.Background(), ip.IP); err != nil {
			return err
		}
	}

	return nil
}
*/
func TestExample_File_Creation(t *testing.T) {

	resource.Test(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		// Providers:    testAccProviders,
		CheckDestroy: testAccCheckFileDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFileExample,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFileExists("example_file.foobar", "foobar"),
				),
			},
		},
	})
}

func testAccCheckFileDestroyed(s *terraform.State) error {

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "example_file" {
			return fmt.Errorf("File still exists")
		}
	}
	return nil
}

func testAccCheckFileExists(n string, path string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		return nil
	}
}

var testAccFileExample = `
resource "example_file" "foobar" {
  path = "./foo"
}`
