package outscale

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/outscale/terraform-provider-outscale/utils"
)

func TestAccOutscaleDSLoadBalancerTags_basic(t *testing.T) {
	t.Parallel()
	r := acctest.RandString(4)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: getTestAccDSODSutscaleOAPILBUDSTagsConfig(r, utils.GetRegion()),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckODSutscaleOAPILBUDSTagsExists("data.outscale_load_balancer_tags.testds"),
					resource.TestCheckResourceAttr(
						"data.outscale_load_balancer_tags.testds", "tags.#", "1"),
				)},
		},
	})
}

func testAccCheckODSutscaleOAPILBUDSTagsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LBU Tag DS ID is set")
		}

		return nil
	}
}

func getTestAccDSODSutscaleOAPILBUDSTagsConfig(r string, zone string) string {
	return fmt.Sprintf(`
		resource "outscale_load_balancer" "barTags" {
			subregion_names    = ["%sa"]
			load_balancer_name = "foobar-terraform-elb-%s"
		
			listeners {
				backend_port           = 8000
				backend_protocol       = "HTTP"
				load_balancer_port     = 80
				load_balancer_protocol = "HTTP"
			}
		
			tags {
				key = "name"
				value = "baz"
			}
		}
		
		data "outscale_load_balancer_tags" "testds" {
			load_balancer_names = [outscale_load_balancer.barTags.id]
		}
	`, zone, r)
}
