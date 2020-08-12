# harbor_project

`harbor_project` is a resource for managing a project.

## Example Usage

```hcl
resource "harbor_project" "myproject" {
  name                  = "myproject"
  public                = true
  auto_scan             = true
  prevent_vulnerability = true
  severity              = "critical"
}
```

## Argument Reference

* `name` - (Required) Project name.
* `public` - (Optional) Set project to be public. Defaults to `false`.
* `auto_scan` - (Optional) Set automatic scan on push. Defaults to `false`.
* `content_trust` - (Optional) Enable content trust. Defaults to `false`.
* `prevent_vulnerability` - (Optional) Prevent image pull on vulnerability. Defaults to `false`.
* `severity` - (Optional) Vulnerability severity level. Defaults to `low`.

## Attribute Reference

* `id` - Project ID.