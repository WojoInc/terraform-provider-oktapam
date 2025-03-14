---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oktapam_kubernetes_cluster_connection Resource - terraform-provider-oktapam"
subcategory: ""
description: |-
  Beta Feature: A set of details describing how to connect to an existing Kubernetes Cluster. NOTE: This is only available if the ASA Team has the Kubernetes Beta feature enabled; contact support@okta.com if you wish to participate in the Beta.
---

# oktapam_kubernetes_cluster_connection (Resource)

Beta Feature: A set of details describing how to connect to an existing Kubernetes Cluster. NOTE: This is only available if the ASA Team has the Kubernetes Beta feature enabled; contact support@okta.com if you wish to participate in the Beta.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_id` (String) The ASA ID of a Kubernetes cluster.

### Optional

- `api_url` (String) The URL to access the control plane of the Kubernetes cluster.
- `public_certificate` (String) The certificate expected when connecting to the Kubernetes cluster.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# Kubernetes Cluster Connection can be imported using ID of this resource, e.g.,
terraform import oktapam_kubernetes_cluster_connection.example {{id}}
```
