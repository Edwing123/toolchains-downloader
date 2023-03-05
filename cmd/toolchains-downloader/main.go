package main

import (
	"fmt"

	"git.edwing123.dev/toolchains-downloader/pkgs/platform"
)

const (
	ZigReleasesJSONURL string = "https://ziglang.org/download/index.json"
)

func main() {
	flags := GetFlags()
	platformInfo := platform.GetInfo()

	fmt.Println(flags)
	fmt.Println(platformInfo)
}
