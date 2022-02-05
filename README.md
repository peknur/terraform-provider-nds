#  Net Data Source Terraform Provider
[![Tests](https://github.com/peknur/terraform-provider-nds/actions/workflows/test.yml/badge.svg)](https://github.com/peknur/terraform-provider-nds/actions/workflows/test.yml)
[![golangci-lint](https://github.com/peknur/terraform-provider-nds/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/peknur/terraform-provider-nds/actions/workflows/golangci-lint.yml)
[![release](https://github.com/peknur/terraform-provider-nds/actions/workflows/release.yml/badge.svg)](https://github.com/peknur/terraform-provider-nds/actions/workflows/release.yml)

## Build provider

Run the following command to build the provider

```shell
$ make build
```

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory. 

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```
