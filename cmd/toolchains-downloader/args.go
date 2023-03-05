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
	// ToolChainKind represents the kind
	// of toolchain to download.
	ToolChainKind ToolChainKind

	// ToolChainsDir represents the directory
	// where the toolchain will be stored.
	ToolChainsDir string
}

// Gets the provided flags.
func GetFlags() Flags {
	toolChainKind := flag.String(
		"toolchain-kind",
		"",
		"The kind of toolchain to download: [zig, go]",
	)

	toolChainsDir := flag.String(
		"toolchains-dir",
		"",
		"the directory where the toolchain will be stored.",
	)

	flag.Parse()

	// Validate provided values.
	if *toolChainKind != "zig" && *toolChainKind != "go" {
		panic("Invalid value for flag toolchain-kind.")
	}

	if strings.TrimSpace(*toolChainsDir) == "" {
		panic("Flag toolchains-dir is required.")
	}

	return Flags{
		ToolChainKind: ToolChainKind(*toolChainKind),
		ToolChainsDir: *toolChainsDir,
	}
}
