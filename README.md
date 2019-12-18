# `terraform-provider-harbor`

A [Terraform][1] plugin for managing [Harbor][2].

## Contents

* [Installation](#installation)
* [`harbor` Provider](#provider-configuration)
* [Resources](#resources)
  * [`harbor_project`](#harbor_project)
* [Requirements](#requirements)

## Installation

Download and extract the [latest
release](https://github.com/Ouest-France/terraform-provider-harbor/releases/latest) to
your [terraform plugin directory][third-party-plugins] (typically `~/.terraform.d/plugins/`)

## Provider Configuration

### Example

Example provider.
```hcl
provider "harbor" {
  address  = "harbor.mycompany.com"
  user     = "myuser"
  password = "mypassword"
}
```

| Property            | Description                | Type    | Required    | Default    |
| ----------------    | -----------------------    | ------- | ----------- | ---------- |
| `address`           | harbor server address      | String  | true        |            |
| `user`              | harbor username            | String  | true        |            |
| `password`          | harbor password            | String  | true        |            |

## Resources
### `harbor_project`

A resource for managing a project.

#### Example

```hcl
resource "harbor_project" "myproject" {
  name                  = "myproject"
  public                = true
  auto_scan             = true
  prevent_vulnerability = true
  severity              = "critical"
}
```

#### Arguments

| Property                | Description                          | Type    | Required    | Default    |
| ----------------        | -----------------------              | ------- | ----------- | ---------- |
| `name`                  | Project name                         | String  | true        |            |
| `public`                | Set registry to be public            | Bool    | false       | `false`    |
| `auto_scan`             | Set automatic scan on push           | Bool    | false       | `false`    |
| `content_trust`         | Enable content trust                 | Bool    | false       | `false`    |
| `prevent_vulnerability` | Prevent image pull on vulnerability  | Bool    | false       | `false`    |
| `severity`              | Vulnerability severity level         | String  | false       | `low`      |

#### Attributes

| Property             | Description                                    |
| ----------------     | -----------------------                        |
| `id`                 | Project ID                                     |

## Requirements
* Harbor >= 1.10.0

[1]: https://www.terraform.io
[2]: https://goharbor.io
