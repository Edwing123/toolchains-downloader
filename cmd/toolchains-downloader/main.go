package main

import (
	"fmt"

	"git.edwing123.dev/toolchains-downloader/pkgs/platform"
)

var (
	// Map with URLs containing releases information + download links.
	ReleaseInfoURLs = map[string]string{
		"zig": "https://ziglang.org/download/index.json",
		"go":  "",
	}
)

func main() {
	flags := GetFlags()
	platformInfo := platform.GetInfo()

	fmt.Println(flags)
	fmt.Println(platformInfo)
}
