---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oktapam_group Data Source - terraform-provider-oktapam"
subcategory: ""
description: |-
  Returns a previously created ASA Group. For details, see Groups https://help.okta.com/asa/en-us/Content/Topics/Adv_Server_Access/docs/setup/groups.htm.
---

# oktapam_group (Data Source)

Returns a previously created ASA Group. For details, see [Groups](https://help.okta.com/asa/en-us/Content/Topics/Adv_Server_Access/docs/setup/groups.htm).



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The human-readable name of the resource. Values are case-sensitive.

### Read-Only

- `deleted_at` (String) The UTC time of resource deletion. Format is '2022-01-01 00:00:00 +0000 UTC'.
- `id` (String) The ID of this resource.
- `roles` (Set of String) A list of roles for the ASA Group. Options are 'access_user', 'access_admin', and 'reporting_user'.


