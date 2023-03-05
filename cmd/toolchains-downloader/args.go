package main

import (
	"flag"
	"strings"
)

// Represents a kind of toolchain.
type ToolChainKind string

const (
	ZigKind ToolChainKind = "zig"
	GoKind  ToolChainKind = "go"
)

// Flags represent the flags passed
// to the program.
type Flags struct {
	// The kind of toolchain to download.
	Kind ToolChainKind

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
		Kind: ToolChainKind(*kind),
		Dir:  *dir,
	}
}
