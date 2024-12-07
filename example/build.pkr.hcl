# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

packer {
  required_plugins {
    synopkg = {
      version = ">=v0.1.0"
      source  = "github.com/hashicorp/synopkg"
    }
  }
}

source "synopkg-my-builder" "foo-example" {
  mock = local.foo
}

source "synopkg-my-builder" "bar-example" {
  mock = local.bar
}

build {
  sources = [
    "source.synopkg-my-builder.foo-example",
  ]

  source "source.synopkg-my-builder.bar-example" {
    name = "bar"
  }

  provisioner "synopkg-my-provisioner" {
    only = ["synopkg-my-builder.foo-example"]
    mock = "foo: ${local.foo}"
  }

  provisioner "synopkg-my-provisioner" {
    only = ["synopkg-my-builder.bar"]
    mock = "bar: ${local.bar}"
  }

  post-processor "synopkg-my-post-processor" {
    mock = "post-processor mock-config"
  }
}
