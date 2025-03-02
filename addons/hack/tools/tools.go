//go:build tools

// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package tools imports things required by build scripts, to force `go mod` to see them as dependencies

package tools

import (
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
)
