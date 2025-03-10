---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oktapam_ad_connections Data Source - terraform-provider-oktapam"
subcategory: ""
description: |-
  A list of ASA AD Connections associated with an ASA Team.
---

# oktapam_ad_connections (Data Source)

A list of ASA AD Connections associated with an ASA Team.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `certificate_id` (String) If `true`, only connections with a matching certificate ID are returned.
- `gateway_id` (String) If `true`, only connections with a matching gateway ID are returned.
- `include_cert_details` (Boolean) If `true`, results also include certificate details

### Read-Only

- `ad_connections` (List of Object) A list of ASA AD Connections associated with an ASA Team. (see [below for nested schema](#nestedatt--ad_connections))
- `id` (String) The ID of this resource.

<a id="nestedatt--ad_connections"></a>
### Nested Schema for `ad_connections`

Read-Only:

- `certificate_id` (String)
- `deleted_at` (String)
- `domain` (String)
- `domain_controllers` (Set of String)
- `gateway_id` (String)
- `id` (String)
- `name` (String)


