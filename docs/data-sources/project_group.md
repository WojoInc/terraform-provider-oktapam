---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oktapam_project_group Data Source - terraform-provider-oktapam"
subcategory: ""
description: |-
  Returns a previously created ASA Group assigned to a given ASA Project. For details, see Add a Group to a Project https://help.okta.com/asa/en-us/Content/Topics/Adv_Server_Access/docs/setup/add-a-group-to-project.htm.
---

# oktapam_project_group (Data Source)

Returns a previously created ASA Group assigned to a given ASA Project. For details, see [Add a Group to a Project](https://help.okta.com/asa/en-us/Content/Topics/Adv_Server_Access/docs/setup/add-a-group-to-project.htm).



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_name` (String) The human-readable name of the ASA Group. Values are case-sensitive.
- `project_name` (String) The human-readable name of the ASA Project. Values are case-sensitive.

### Read-Only

- `create_server_group` (Boolean) If `true`, 'sftd' (ASA Server Agent) creates a corresponding local (unix or windows) group in the ASA Project's servers.
- `deleted_at` (String) The UTC time of resource deletion. Format is '2022-01-01 00:00:00 +0000 UTC'.
- `id` (String) The ID of this resource.
- `removed_at` (String) UTC time of resource removal from parent resource. Format is '2022-01-01 00:00:00 +0000 UTC'.
- `server_access` (Boolean) If `true`, members of this ASA Group have access to the ASA Project servers.
- `server_admin` (Boolean) If `true`, members of ASA Group have sudo permissions on ASA Project servers.
- `servers_selector` (Map of String) Enables access to ASA Servers with labels matching all selectors. Only available to customers that have the Early Access Policy Sync feature enabled on their team.


