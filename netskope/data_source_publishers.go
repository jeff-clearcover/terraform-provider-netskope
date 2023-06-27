package netskope

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jeff-clearcover/netskope-api-client-go/nsgo"
)

func dataSourcePublishersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Collect Diags
	var diags diag.Diagnostics

	//Set Filter
	filter := d.Get("filter").(string)

	//Init a client instance
	nsclient := m.(*nsgo.Client)

	//Get Publishers
	pubs, err := nsclient.GetPublishersWithFilter(filter)
	if err != nil {
		return diag.FromErr(err)
	}

	jsonData, _ := json.Marshal(pubs)

	pubsStruct := nsgo.PublishersList{}
	json.Unmarshal(jsonData, &pubsStruct)

	newjsonData, _ := json.Marshal(pubsStruct.Publishers)
	pubsMap := make([]map[string]interface{}, 0)
	json.Unmarshal(newjsonData, &pubsMap)

	if err := d.Set("publishers", pubsMap); err != nil {
		/*
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create HashiCups client",
				Detail:   string([]byte(newjsonData)),
			})
			return diags
		*/
		return diag.FromErr(err)
	}

	/*
		//Get Publisher
		pubs, err := nsclient.GetPublishers()
		if err != nil {
			return diag.FromErr(err)
		}

		pubsMap := make([]map[string]interface{}, 0)
		err = json.NewDecoder(pubs).Decode(&pubsMap)
		if err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("publishers", pubsMap); err != nil {
			return diag.FromErr(err)
		}

		// always run
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	*/
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags

}

func dataSourcePublishers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePublishersRead,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"publishers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"publisher_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"publisher_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"publisher_upgrade_profiles_external_id": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"upgrade_failed_reason": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"detail": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"error_code": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"timestamp": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"lbrokerconnect": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"upgrade_request": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"common_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"registered": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"upgrade_status": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status_failure_code": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"upstat": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": {
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Computed: true,
						},
						"stitcher_id": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"assessment": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"eee_support": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"hdd_free": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"hdd_total": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"latency": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
