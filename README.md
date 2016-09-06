# terraform-provider-pikacloud
Terraform provider for https://pikacloud.com

[![Build Status](https://travis-ci.org/pikacloud/terraform-provider-pikacloud.svg?branch=master)](https://travis-ci.org/pikacloud/terraform-provider-pikacloud)

## Installing

If you have Go installed, you can just do:
```
go get github.com/pikacloud/terraform-provider-pikacloud
```
This will automatically download, compile and install the app.

After that you should have a `terraform-provider-pikacloud` executable in your `$GOPATH/bin`.

## Provider configuration

```
provider "pikacloud" {
  token = "PIKACLOUD_API_TOKEN"
}
```

## Resources

### DNS Zones

```
resource "pikacloud_zone" "example" {
  domain_name "example.com"
}
```
#### Argument Reference

- `domain_name` - (Required) Domain name

#### Attributes Reference

- `id` - The ID of the zone
- `serial` - Zone serial
