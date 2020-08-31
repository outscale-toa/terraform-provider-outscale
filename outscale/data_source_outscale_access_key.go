package outscale

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/antihax/optional"
	oscgo "github.com/marinsalinas/osc-sdk-go"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceOutscaleAccessKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOutscaleAccessKeyRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"access_key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"creation_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_modification_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ACTIVE", "INACTIVE"}, false),
			},
			"request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceOutscaleAccessKeyRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OSCAPI

	filters, filtersOk := d.GetOk("filter")
	accessKeyID, accessKeyOk := d.GetOk("access_key_id")
	state, stateOk := d.GetOk("state")

	if !filtersOk && !accessKeyOk && !stateOk {
		return fmt.Errorf("One of filters, access_key_id or state must be assigned")
	}

	filterReq := &oscgo.FiltersAccessKeys{}
	if filtersOk {
		filterReq = buildOutscaleDataSourceAccessKeyFilters(filters.(*schema.Set))
	}
	if accessKeyOk {
		filterReq.SetAccessKeyIds([]string{accessKeyID.(string)})
	}
	if stateOk {
		filterReq.SetStates([]string{state.(string)})
	}

	var resp oscgo.ReadAccessKeysResponse
	var err error
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		resp, _, err = conn.AccessKeyApi.ReadAccessKeys(context.Background(), &oscgo.ReadAccessKeysOpts{
			ReadAccessKeysRequest: optional.NewInterface(oscgo.ReadAccessKeysRequest{
				Filters: filterReq,
			}),
		})
		if err != nil {
			if strings.Contains(fmt.Sprint(err), "RequestLimitExceeded:") {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return err
	}

	if len(resp.GetAccessKeys()) == 0 {
		return fmt.Errorf("Unable to find Access Key")
	}

	if len(resp.GetAccessKeys()) > 1 {
		return fmt.Errorf("multiple results returned, please use a more specific criteria in your query")
	}

	accessKey := resp.GetAccessKeys()[0]

	if err := d.Set("access_key_id", accessKey.GetAccessKeyId()); err != nil {
		return err
	}
	if err := d.Set("creation_date", accessKey.GetCreationDate()); err != nil {
		return err
	}
	if err := d.Set("last_modification_date", accessKey.GetLastModificationDate()); err != nil {
		return err
	}
	if err := d.Set("state", accessKey.GetState()); err != nil {
		return err
	}
	if err := d.Set("request_id", resp.ResponseContext.GetRequestId()); err != nil {
		return err
	}

	d.SetId(accessKey.GetAccessKeyId())

	return nil
}

func buildOutscaleDataSourceAccessKeyFilters(set *schema.Set) *oscgo.FiltersAccessKeys {
	var filters oscgo.FiltersAccessKeys
	for _, v := range set.List() {
		m := v.(map[string]interface{})
		var filterValues []string
		for _, e := range m["values"].([]interface{}) {
			filterValues = append(filterValues, e.(string))
		}

		switch name := m["name"].(string); name {
		case "access_key_ids":
			filters.SetAccessKeyIds(filterValues)
		case "states":
			filters.SetStates(filterValues)
		default:
			log.Printf("[Debug] Unknown Filter Name: %s.", name)
		}
	}
	return &filters
}