---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scalingo_autoscaler Resource - terraform-provider-scalingo"
subcategory: ""
description: |-
  Resource representing an autoscaler of an application, setting rules to automatically scale up and down containers of the app
---

# scalingo_autoscaler (Resource)

Resource representing an autoscaler of an application, setting rules to automatically scale up and down containers of the app

## Example Usage

```terraform
resource "scalingo_app" "test_app" {
  name = "terraform-test-autoscaler"
}

# Create an autoscaler to scale 'web' containers to ensure CPU consumption stays under 80%
resource "scalingo_autoscaler" "test_autoscaler" {
  app            = scalingo_app.test_app.id
  container_type = "web"
  min_containers = 2
  max_containers = 10
  metric         = "cpu"
  target          = 0.8
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app` (String) ID or Name of the targeted application
- `container_type` (String) Container type targeted by the autoscaler (web, worker, etc.)
- `max_containers` (Number) Maximum number of containers (autoscaler won't get over it)
- `metric` (String) Watched metric to base the autoscaling on (cpu, ram, etc.)
- `min_containers` (Number) Minimum number of containers (autoscaler won't get under it)
- `target` (Number) Target reference value to base the autoscaling algorithm on

### Optional

- `disabled` (Boolean) Disable without deleting the autoscaler

### Read-Only

- `id` (String) The ID of this resource.

