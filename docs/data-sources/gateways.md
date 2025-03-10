---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oktapam_gateways Data Source - terraform-provider-oktapam"
subcategory: ""
description: |-
  Returns a list of all ASA Gateways connected to the ASA Team specified in the OKTAPAMTEAM environment variable. For details, see [Advanced Access Server Gateways](https://help.okta.com/asa/en-us/Content/Topics/AdvServer_Access/docs/gateways-and-bastions.htm) on ASA Gateways and bastions.
---

# oktapam_gateways (Data Source)

Returns a list of all ASA Gateways connected to the ASA Team specified in the OKTAPAM_TEAM environment variable. For details, see [Advanced Access Server Gateways](https://help.okta.com/asa/en-us/Content/Topics/Adv_Server_Access/docs/gateways-and-bastions.htm) on ASA Gateways and bastions.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `contains` (String) If a value is provided, the results are filtered to only contain resources that contain the value in the name field.

### Read-Only

- `gateways` (List of Object) Returns a list of all ASA Gateways connected to the ASA Team specified in the OKTAPAM_TEAM environment variable. For details, see [Advanced Access Server Gateways](https://help.okta.com/asa/en-us/Content/Topics/Adv_Server_Access/docs/gateways-and-bastions.htm) on ASA Gateways and bastions. (see [below for nested schema](#nestedatt--gateways))
- `id` (String) The ID of this resource.

<a id="nestedatt--gateways"></a>
### Nested Schema for `gateways`

Read-Only:

- `access_address` (String)
- `cloud_provider` (String)
- `default_address` (String)
- `description` (String)
- `id` (String)
- `labels` (Map of String)
- `name` (String)
- `refuse_connections` (Boolean)


