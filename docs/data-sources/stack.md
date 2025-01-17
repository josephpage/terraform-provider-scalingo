---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scalingo_stack Data Source - terraform-provider-scalingo"
subcategory: ""
description: |-
  List of available stacks to base applications on
---

# scalingo_stack (Data Source)

List of available stacks to base applications on

## Example Usage

```terraform
data "scalingo_stack" "my_chosen_stack" {
  name = "scalingo-22"
}

resource "scalingo_app" "my_test_app" {
  name = "terraform-testapp"

  stack_id = data.scalingo_stack.my_chosen_stack.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Slug name of the stack (scalingo-18, scalingo-20, …)

### Read-Only

- `base_image` (String) Base docker image on which is based the stack
- `default` (Boolean) Is it the current default stack?
- `deprecated_at` (String) When has been/will be deprecated the stack
- `description` (String) Textual description of the stack
- `id` (String) The ID of this resource.


