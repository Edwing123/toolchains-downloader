package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"git.edwing123.dev/toolchains-downloader/pkgs/platform"
)

var (
	// Map with URLs containing releases information + download links.
	ReleaseInfoURLs = map[string]string{
		"zig": "https://ziglang.org/download/index.json",
		"go":  "",
	}
)

// Fetches the JSON at the provided url using the provided client.
func FetchJSON(url string, client *http.Client) (map[string]any, error) {
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data map[string]any

	jsonDecoder := json.NewDecoder(res.Body)

	err = jsonDecoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Creates a new HTTP client.
func NewClient() *http.Client {
	client := http.Client{}
	return &client
}

// Displays a banner with information.
func DisplayBanner(platformInfo platform.Info, flags Flags) {
	fmt.Println("== Information")
	fmt.Println("Toolchain kind:", flags.Kind)
	fmt.Println("Output directory:", flags.Dir)
	fmt.Println("OS:", platformInfo.Os)
	fmt.Println("CPU Architecture:", platformInfo.CPUArch)
}

// Fetches the file from url using client.
func FetchTarball(url string, client *http.Client) ([]byte, error) {
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetZigTarballURLAndSize(releasesInfo map[string]any, platformInfo platform.Info) (string, string) {
	masterMap := releasesInfo["master"].(map[string]any)

	cpu := "x86_64"
	targetField := fmt.Sprintf("%s-%s", cpu, platformInfo.Os)

	targetInfo := masterMap[targetField].(map[string]any)

	tarballURL := targetInfo["tarball"].(string)
	tarballSize := targetInfo["size"].(string)
	return tarballURL, tarballSize
}

// Decompress the tarball file using the system "tar" command.
func DecompressTarball(file string, dir string) error {
	command := exec.Command("tar", "-C", dir, "-xf", file)
	return command.Run()
}

func main() {
	flags := GetFlags()
	platformInfo := platform.GetInfo()
	client := NewClient()

	// Display informative banner.
	DisplayBanner(platformInfo, flags)

	// Get releases info based on toolchain kind.
	releasesInfoURL := ReleaseInfoURLs[flags.Kind]

	fmt.Println()
	fmt.Println("== Getting releases information")
	fmt.Println("Using URL:", releasesInfoURL)

	releasesInfo, err := FetchJSON(releasesInfoURL, client)
	if err != nil {
		fmt.Println("The following error ocurred while getting release information:", err)
		os.Exit(0)
	}

	// Get the URL of the tarball file with the latest compiler toolchain.
	fmt.Println()
	fmt.Println("== Getting tarball information")

	var tarballURL string
	var tarballSize string

	switch flags.Kind {
	case "zig":
		tarballURL, tarballSize = GetZigTarballURLAndSize(releasesInfo, platformInfo)
	case "go":
		fmt.Println("Downloading Go is not supported yet.")
		os.Exit(0)
	}

	fmt.Println("Tarball URL:", tarballURL)
	fmt.Println("Tarball size:", tarballSize)

	// Get tarball contents.
	fmt.Println()
	fmt.Println("== Downloading tarball file")

	tarball, err := FetchTarball(tarballURL, client)

	if err != nil {
		fmt.Println("The following error ocurred while downloading the file:", err)
		os.Exit(0)
	}

	fmt.Println("Done")

	// Store tarball contents to temporary file.
	fmt.Println()
	fmt.Println("== Storing tarball file to temporary location")

	tmpTarballFile, err := os.CreateTemp("", flags.Kind)
	if err != nil {
		fmt.Println("The following error ocurred while creating temporary file:", err)
		os.Exit(0)
	}
	defer tmpTarballFile.Close()

	// Write to file.
	_, err = tmpTarballFile.Write(tarball)
	if err != nil {
		fmt.Println("The following error ocurred writing to temporary file:", err)
		os.Exit(0)
	}

	fmt.Println("Done")

	// Decompress tarball to output directory.
	fmt.Println()
	fmt.Println("== Decompressing tarball into output directory")

	tmpTarballFileName := tmpTarballFile.Name()
	fmt.Println("tarball file:", tmpTarballFileName)

	err = DecompressTarball(tmpTarballFileName, flags.Dir)
	if err != nil {
		fmt.Println("The following error ocurred decompressing file:", err)
		os.Exit(0)
	}

	fmt.Println("Done")
}
