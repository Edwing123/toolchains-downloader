package platform

import "go/build"

// Info constains information about the computer platform.
type Info struct {
	Os      string
	CPUArch string
}

// Returns the information of the computer platform.
func GetInfo() Info {
	context := build.Default

	return Info{
		Os:      context.GOOS,
		CPUArch: context.GOARCH,
	}
}
