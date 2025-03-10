---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oktapam_server_enrollment_token Resource - terraform-provider-oktapam"
subcategory: ""
description: |-
  A token used to enroll servers in an ASA Project. For details, see Enroll a server https://help.okta.com/asa/en-us/Content/Topics/Adv_Server_Access/docs/setup/enrolling-a-server.htm.
---

# oktapam_server_enrollment_token (Resource)

A token used to enroll servers in an ASA Project. For details, see [Enroll a server](https://help.okta.com/asa/en-us/Content/Topics/Adv_Server_Access/docs/setup/enrolling-a-server.htm).



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) If a value is provided, the results are filtered to only contain resources whose description contains that value.
- `project_name` (String) The human-readable name of the ASA Project. Values are case-sensitive.

### Read-Only

- `created_by_user` (String) The ASA User that created the resource.
- `id` (String) The ID of this resource.
- `issued_at` (String) The UTC time when the token was issued. Format is '2022-01-01 00:00:00 +0000 UTC'.
- `token` (String) The secret used for resource enrollment.

## Import

Import is supported using the following syntax:

```shell
# Server Enrollment Token can be imported using Project Name and ID of the resource separated by a forward slash (/), e.g.,
terraform import oktapam_server_enrollment_token.example {{project_name}}/{{id}}
```
