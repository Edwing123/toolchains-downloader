package main

import (
	"fmt"

	"git.edwing123.dev/toolchains-downloader/pkgs/platform"
)

const (
	// URL of the JSON file with Zig releases information.
	ZigReleasesURL string = "https://ziglang.org/download/index.json"
)

func main() {
	flags := GetFlags()
	platformInfo := platform.GetInfo()

	fmt.Println(flags)
	fmt.Println(platformInfo)
}
