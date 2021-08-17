package outscale

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceOutscaleOAPIFlexibleGpuCatalog_basic(t *testing.T) {
	region := fmt.Sprintf("%sa", os.Getenv("OUTSCALE_REGION"))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceOutscaleOAPIFlexibleGpuCatalogConfig(region),
			},
		},
	})
}

func testAccDataSourceOutscaleOAPIFlexibleGpuCatalogConfig(region string) string {
	return fmt.Sprintf(`
		data "outscale_flexible_gpu_catalog" "catalog-fGPU" { }`)
}
