package main

import (
	"flag"
	"strings"
)

// Flags represent the flags passed
// to the program.
type Flags struct {
	// The kind of toolchain to download.
	Kind string

	// The directory where the toolchain will be stored.
	Dir string
}

// Gets the provided flags.
func GetFlags() Flags {
	kind := flag.String(
		"kind",
		"",
		"The toolchain to download: [zig, go]",
	)

	dir := flag.String(
		"dir",
		"",
		"the directory where the toolchain will be stored.",
	)

	flag.Parse()

	// Validate provided values.
	if *kind != "zig" && *kind != "go" {
		panic("Invalid value for flag kind.")
	}

	if strings.TrimSpace(*dir) == "" {
		panic("Flag dir is required.")
	}

	return Flags{
		Kind: *kind,
		Dir:  *dir,
	}
}
