# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

locals {
  foo = data.synopkg-my-datasource.mock-data.foo
  bar = data.synopkg-my-datasource.mock-data.bar
}