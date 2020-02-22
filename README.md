Dummy Terraform Provider
==================

## Motivation 
I wanted to learn how to build a custom terraform provider. So I created a dummy provider that can create, update and delete empty files.

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.11 

Building The Provider
---------------------

Clone repository: 
```
$ git clone git@github.com:jackap/terraform-provider-tfsystem
```

Enter the provider directory and build the provider

```sh
$ cd terraform-provider-tfsystem
$ go build
```
Copy the generated binary inside your terraform project:

```sh
$ mkdir my-terraform-project
$ mv ../terraform-provider-tfsystem .terraform/plugins/<OS>/
```

Sample configuration file
----------------------

```terraform
provider "tfsystem"{}

resource "tfsystem_file" "foo" {
    path = "./bar"
}
```