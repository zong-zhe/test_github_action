// Copyright 2021 The KCL Authors. All rights reserved.

package opt

// User input options from the cli.
type Options struct {
	init *InitOptions
}

// Input options of 'kpm init'.
type InitOptions struct {
	Name     string
	InitPath string
}
