---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_infra_join_token Resource - terraform-provider-bloxone"
subcategory: ""
description: |-
  
---

# bloxone_infra_join_token (Resource)



## Example Usage

```terraform
resource "time_offset" "one_week" {
  offset_days = 7
}

resource "bloxone_infra_join_token" "example" {
  name = "example_join_token"

  # Other optional fields
  description = "Join token for test site"
  expires_at  = time_offset.one_week.rfc3339
  tags = {
    site = "Test Site"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `description` (String)
- `expires_at` (String)
- `tags` (Map of String)

### Read-Only

- `deleted_at` (String)
- `id` (String) The resource identifier.
- `join_token` (String, Sensitive)
- `last_used_at` (String)
- `status` (String)
- `token_id` (String) first half of the token.
- `use_counter` (Number)