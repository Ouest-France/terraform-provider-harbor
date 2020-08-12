# Harbor Provider

The Harbor provider is used to interact with the resources supported by [Harbor](https://goharbor.io/) API.

## Example Usage

```hcl
# Provider configuration
terraform {
  required_providers {
    harbor = {
      source  = "Ouest-France/harbor"
    }
  }
}

provider "harbor" {
  address  = "harbor.mycompany.com"
  user     = "myuser"
  password = "mypassword"
}

...
```

## Argument Reference

* `address` - (Required) Harbor server address formatted like `harbor.mycompany.com`.
* `user` - (Required) Harbor user to access the API.
* `password` - (Required) Harbor password to access the API.

## Requirements

* Harbor >= 1.10.0