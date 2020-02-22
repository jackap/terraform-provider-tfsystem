package tfsystem

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestExample_File_Creation(t *testing.T) {

	resource.Test(t, resource.TestCase{
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
