# terraform-provider-pikacloud
Terraform provider for [Pikacloud](https://pikacloud.com) services.

[![Build Status](https://travis-ci.org/pikacloud/terraform-provider-pikacloud.svg?branch=master)](https://travis-ci.org/pikacloud/terraform-provider-pikacloud) [![Slack Chat](https://img.shields.io/badge/slack-join%20the%20chat-00B9FF.svg)](https://slack-invite.pikacloud.com)


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
resource "pikacloud_zone" "example_com" {
  domain_name = "example.com"
}
```
#### Argument Reference

- `domain_name` - (Required) Domain name

#### Attributes Reference

- `id` - The ID of the zone
- `serial` - Zone serial

### DNS Zone records

```
resource "pikacloud_zonerecord" "www_example_com" {
  zone = "${pikacloud_zone.example_com.id}"
  rtype = "A"
  name = "www"
  ipv4 = "${digitalocean_droplet.my_server.ipv4_address}"
}
```
#### Argument Reference

- `rtype` - (Require) Record type, one of A, CNAME, TXT, NS, ALIAS, MX
- `name` - (Optional) Subdomain, can be leave empty to aim zone root (also known as the zone APEX)
- `hostname` - (Required for CNAME, NS, ALIAS or MX) Target hostname
- `ipv4` - (Required for A type) IPv4 address
- `content` - (Required for TXT) Text field
- `priority` - (Required for MX) Priority
- `ttl` - (Optional) Time to live (default: 1800 seconds)

#### Attributes Reference

- `id` - The ID of the zone record
- `zone` - The ID of the zone

## Building from source

- Install Go on your machine
- Set up Gopath
- `git clone` this repository into `$GOPATH/src/github.com/pikacloud/terraform-provider-pikacloud`
- Run `go get` to get dependencies
- Run `go install` to build the binary. You will now find the binary at `$GOPATH/bin/terraform-provider-pikacloud.`

## Running tests

- Get an API Token from [Pikacloud](https://pikacloud.com)
- Run tests based on real resources:
```
PIKACLOUD_TOKEN=<api token> TF_ACC=1 go test ./... -v
```
