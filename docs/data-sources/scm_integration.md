---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scalingo_scm_integration Data Source - terraform-provider-scalingo"
subcategory: ""
description: |-
  SCM Integrations which are attached to an account, required to use github/gitlab integrations with SCM Repo Links
---

# scalingo_scm_integration (Data Source)

SCM Integrations which are attached to an account, required to use github/gitlab integrations with SCM Repo Links

## Example Usage

```terraform
data "scalingo_scm_integration" "github" {
  scm_type = "github"
}

resource "scalingo_app" "test_app" {
  name = "terraform-test-scm"
}

resource "scalingo_scm_repo_link" "github-link" {
  auth_integration_uuid = data.scalingo_scm_integration.github.id
  app                   = scalingo_app.test_app.id
  source                = "https://github.com/nalpskc/sample-node-express"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `scm_type` (String) Type of SCM integration (github/gitlab)
- `url` (String) URL to the SCM integration provider

### Read-Only

- `avatar_url` (String) Avatar URL from the integration platform account
- `email` (String) email from the integration platform account
- `id` (String) The ID of this resource.
- `owner_id` (String) ID of the user owning the integration (self when doing the request)
- `profile_url` (String) Profile URL from the integration platform account
- `uid` (String) Unique identifier of the SCM integration
- `username` (String) Username of the integration platform account


