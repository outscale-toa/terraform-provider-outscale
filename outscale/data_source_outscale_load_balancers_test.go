package outscale

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-outscale/utils"
)

func TestAccOutscaleOAPIDSLBSU_basic(t *testing.T) {
	t.Parallel()
	numLbu := acctest.RandIntRange(0, 50)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: "outscale_load_balancer.bar",
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckOutscaleOAPILBUDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDSOutscaleOAPILBsUConfig(utils.GetRegion(), numLbu),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.outscale_load_balancers.test", "load_balancer.#", "1"),
				)},
		},
	})
}

func testAccDSOutscaleOAPILBsUConfig(zone string, numLbu int) string {
	return fmt.Sprintf(`
	resource "outscale_load_balancer" "bar" {
		subregion_names = ["%sa"]
		load_balancer_name        = "foobar-terraform-elb%d"

		listeners {
			backend_port      = 8000
			backend_protocol  = "HTTP"
			load_balancer_port = 80

			// Protocol should be case insensitive
			load_balancer_protocol = "HTTP"
		}

		tags {
			key = "name"
			value = "baz"
		}
	}

	data "outscale_load_balancers" "test" {
		load_balancer_name = ["${outscale_load_balancer.bar.id}"]
	}
`, zone, numLbu)
}
