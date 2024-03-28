package outscale

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oscgo "github.com/outscale/osc-sdk-go/v2"
	"github.com/outscale/terraform-provider-outscale/utils"
)

func TestAccOthers_ServerCertificate_basic(t *testing.T) {
	t.Parallel()
	resourceName := "outscale_server_certificate.test"
	rName := acctest.RandomWithPrefix("acc-test")
	rNameUpdated := acctest.RandomWithPrefix("acc-test")
	body := `-----BEGIN CERTIFICATE-----
MIIFETCCAvkCFE1QlISgW8h5/akhNlZzb+or8HgYMA0GCSqGSIb3DQEBCwUAMEUx
CzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRl
cm5ldCBXaWRnaXRzIFB0eSBMdGQwHhcNMjEwOTA5MDAyNDQzWhcNMzEwOTA3MDAy
NDQzWjBFMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UE
CgwYSW50ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIICIjANBgkqhkiG9w0BAQEFAAOC
Ag8AMIICCgKCAgEA04T7eUNzkdtMDlcWpzgu/xSg98ZvhnDB75s528OXgw08UFM9
GpqGYu12xbt02z0PJmBBJPy7qm3Dp2OtLuDcATL6QJTi0KX5xvuLTmBnuSoHe5Ec
bMOx2YHR6Hi9hUtl6M4E2y4Lxc5TQonIHtBprTWb4B5SFVdjgh8PFPHWJ8/YCM53
RoUOsPnGb61cfStsK0zd3iUXOiIDDrN8EaMZi4ykqdBL2eztKWf6DSUTzK6sLZJu
MHK8Q5psryKO68gUx8pbwHGSEf+UxhsC1Bik+HJBH8OvULvB8OpBP4IixI0RHT3K
atrv6z+gl3X7606jaRxs4hlKrfORGK6SB6xloi+ptfXSKB1+6L0I2gjcIbpNvM+h
Dsvp1oly4H/QGA/g24dQgVuoB4hetZ5U+IQU231CJAOVLrw958ln746KN5wtQ8ur
IFkMRKb2X4ccAmLQroOxZGxv9MgvpQIJCdC18/fkZdEMKTGK6TmrcvPZ9A8tLYzn
yQcOvJ4r6Yh5Z6TUJWSjPqMPYZQgb2NBSGP8MBuQ67BHMBSO5F+wM+PPURw3clwl
qVJHXN1OLakmwxY2jsrZRxX8P2Q9Nkd/vRWmA6z4TlzEQxaKcJUHgsjbW5sl/FJ1
h/E4CUfCReBXSF2ByWRcMkkAtcWQhzBMWJzAsNUltK4HgT8yY2CUMVSwTXcCAwEA
ATANBgkqhkiG9w0BAQsFAAOCAgEAiDGXtRVWfMvytjRmG7OeDuPvnKEcNSXyUARj
o4vq6z3SOVpfv4gG98p0aT3Xoi3yyWW+zATySOfgVTQkDVtHsBy98k2Z6/jP4Bsi
ryOuxhz4JR9hCyyNq3/RlIKcdTpcKIY0MC+xfVSyAXJXvX+MuGOPqB7/AF842Pyv
vn0Vfkk3K7gMTDITgp7+XVfjX+PF/pjeJ8p4yrSmhUAylPat1oNWtOiEDV/OMaQD
PGAFSEU7+07jz6fQSXjXE8u+uWZ9CUP3F5aRws0kold65aAamfCNXeckEuvsrIrB
hCAXFh06z/amaxR3Sg1pxxx3QNuwRp7KP79eu6Y+1erIwKYqfo/T7lEJn7i+CHJQ
M0veJi5Saoe1mZpq6/fBSnaEtDOw/yQ8+eMe37fiB/jdYb2L3FabkwUJH6Do3BkN
V4WVjjwuT0MHPLJUrsoKDd37GkDsrbhrp/MbgITy1bZWdqaYf0fzDCy+cyUwpWLf
T4ELtBlg7+Hl+1ysiPU2voq1COBiz5MGHNbVj3FI0SteOjFZXMcvfhQ//bksk5Kb
8hpaM5fMheQj6iaLhsW6u4qIJUd4hclrACMwZUIFOSlXs6HZBkmUfB/4XBn5Vas3
zz2wyUFLztD3nO3US8tPSz/I5kXWGpOrPt+UcUPQEGzu5WQ+ZOeD+mQQMCVv0wiz
nmS35ug=
-----END CERTIFICATE-----
`
	private := `-----BEGIN PRIVATE KEY-----
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
0UFH3oBALPqPwDIt0F+wjSaY2bcmCjo=
-----END PRIVATE KEY-----`

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOutscaleServerCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOutscaleOAPIServerCertificateConfig(rName, body, private),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleServerCertificateExists(resourceName),
				),
			},
			{
				Config: testAccOutscaleOAPIServerCertificateConfig(rNameUpdated, body, private),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleServerCertificateExists(resourceName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"request_id", "body", "private_key"},
			},
		},
	})
}

func testAccCheckOutscaleServerCertificateExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := testAccProvider.Meta().(*OutscaleClient).OSCAPI

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}
		exists := false
		var resp oscgo.ReadServerCertificatesResponse
		err := resource.Retry(3*time.Minute, func() *resource.RetryError {
			rp, httpResp, err := conn.ServerCertificateApi.ReadServerCertificates(context.Background()).ReadServerCertificatesRequest(oscgo.ReadServerCertificatesRequest{}).Execute()
			if err != nil {
				return utils.CheckThrottling(httpResp, err)
			}
			resp = rp
			return nil
		})

		if err != nil || len(resp.GetServerCertificates()) == 0 {
			return fmt.Errorf("Server Certificate not found (%s)", rs.Primary.ID)
		}

		for _, server := range resp.GetServerCertificates() {
			if server.GetId() == rs.Primary.ID {
				exists = true
			}
		}

		if !exists {
			return fmt.Errorf("Server Certificate not found (%s)", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckOutscaleServerCertificateDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*OutscaleClient).OSCAPI

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "outscale_server_certificate_link" {
			continue
		}

		exists := false

		var resp oscgo.ReadServerCertificatesResponse
		var err error
		err = resource.Retry(3*time.Minute, func() *resource.RetryError {
			rp, httpResp, err := conn.ServerCertificateApi.ReadServerCertificates(context.Background()).ReadServerCertificatesRequest(oscgo.ReadServerCertificatesRequest{}).Execute()
			if err != nil {
				return utils.CheckThrottling(httpResp, err)
			}
			resp = rp
			return nil
		})

		if err != nil {
			return fmt.Errorf("Server Certificate reading (%s)", rs.Primary.ID)
		}

		for _, server := range resp.GetServerCertificates() {
			if server.GetId() == rs.Primary.ID {
				exists = true
			}
		}

		if exists {
			return fmt.Errorf("Server Certificate still exists (%s)", rs.Primary.ID)
		}
	}
	return nil
}

func testAccOutscaleOAPIServerCertificateConfig(name, body, privateKey string) string {
	return fmt.Sprintf(`
resource "outscale_server_certificate" "test" { 
   name        =  %[1]q
   body        =  %[2]q
   private_key =  %[3]q
}
	`, name, body, privateKey)
}
