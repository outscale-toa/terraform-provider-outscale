package outscale

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-outscale/utils"
)

func TestAccDataSourceOutscaleOAPIFlexibleGpu_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceOutscaleOAPIFlexibleGpuConfig(utils.GetRegion()),
			},
		},
	})
}

func testAccDataSourceOutscaleOAPIFlexibleGpuConfig(region string) string {
	return fmt.Sprintf(`
                resource "outscale_flexible_gpu" "fGPU-1" { 
                        model_name             =  "nvidia-k2"
                        generation             =  "v3"
                        subregion_name         =  "%[1]sa"
                        delete_on_vm_deletion  =   true
                }

		data "outscale_flexible_gpu" "data_fGPU-1" {
			filter {
				name = "flexible_gpu_ids"
				values = [outscale_flexible_gpu.fGPU-1.flexible_gpu_id]
			}
                        filter {
                                name = "delete_on_vm_deletion"
                                values = [true]
                        }
                        filter {
                                name = "generations"
                                values = ["v3"]
                        }
                        filter {
                                name = "states"
                                values = ["allocated"]
                        }
                        filter {
                                name = "model_names"
                                values = ["nvidia-k2"]
                        }
	                filter {
                                name = "subregion_names" 
                                values = ["%[1]s"]
                        }
		}
	`, region)
}
