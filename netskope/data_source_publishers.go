package netskope

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jeff-clearcover/netskope-api-client-go/nsgo"
)

//func flattenPublishersReadData(publishers *[]map[string]interface{}) []interface{} {
//	if publishers != nil {
//		pubs = make([]interface{}, len(*publishers), len(*publishers))
//		for i, publisher := range *publishers {
//			publisher["assessment"] = publishers[i]
//		}
//	}
//}

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
	err = json.Unmarshal(jsonData, &pubsStruct)
	if err != nil {
		return diag.FromErr(err)
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "PUBSSTRUCT",
		Detail:   "PUBSSTRUCT: " + fmt.Sprintf("%+v\n%#v\n", pubsStruct, pubsStruct),
	})

	newjsonData, _ := json.Marshal(pubsStruct.Publishers)

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "NEWJSON",
		Detail:   "NEWJSON: " + fmt.Sprintf("%s", newjsonData),
	})

	pubsMap := make([]map[string]interface{}, 0)
	err = json.Unmarshal(newjsonData, &pubsMap)

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "PUBSMAPERROR",
		Detail:   "PUBSMAP: " + fmt.Sprintf("%+v\n%#v\n", pubsMap, pubsMap),
	})

	//for _, publisher := range pubsMap {
	//	publisher["assessment"] = publisher["assessment"].(*schema.Set)
	//	diags = append(diags, diag.Diagnostic{
	//		Severity: diag.Warning,
	//		Summary:  "PUBLISHER",
	//		Detail:   "PUBLISHER: " + fmt.Sprintf("%s", publisher["assessment"]),
	//	})
	//	return diags
	//}

	if err != nil {
		return diag.FromErr(err)
	}

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
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"timestamp": {
										Type:     schema.TypeFloat,
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
							Type:     schema.TypeList,
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
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Computed: true,
						},
						"stitcher_id": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"assessment": {
							Type:     schema.TypeList,
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
