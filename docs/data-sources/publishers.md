---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netskope_publishers Data Source - terraform-provider-netskope"
subcategory: ""
description: |-
  
---

# netskope_publishers (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) The ID of this resource.

### Read-Only

- `publishers` (List of Object) (see [below for nested schema](#nestedatt--publishers))

<a id="nestedatt--publishers"></a>
### Nested Schema for `publishers`

Read-Only:

- `assessment` (Map of String)
- `common_name` (String)
- `publisher_id` (Number)
- `publisher_name` (String)
- `publisher_upgrade_profiles_id` (String)
- `registered` (Boolean)
- `status` (String)
- `stitcher_id` (Number)
- `upgrade_failed_reason` (String)
- `upgrade_request` (Boolean)

