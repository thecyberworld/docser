package main

import (
	"docser/internal/scanner"
	"docser/internal/upgrade"
	"flag"
	"fmt"
	"os"
)

const (
	owner          = "thecyberworld"
	repo           = "docser"
	currentVersion = "v0.1.0"
)

func main() {

	pRepoLocation := flag.String("d", "", "Directory to be scanned. (Default is current directory)")
	pConfigFile := flag.String("c", "", "Docser config file (Must end in .toml)")
	pUpgrade := flag.Bool("upgrade", false, "Upgrade Docser to latest version")
	showHelp := flag.Bool("h", false, "Displays help menu")

	flag.Parse()

	if *showHelp {
		showHelpMenu()
		os.Exit(0)
	}

	if *pUpgrade {
		upgrade.Start(currentVersion, owner, repo)
		return
	}

	printBanner()
	repositoryPath := *pRepoLocation
	configFile := *pConfigFile
	scanner.ParseConfigAndInitiateScan(configFile, repositoryPath)
}

func showHelpMenu() {
	fmt.Println("Usage: docser -d /path/to/directory -c /path/to/.docser.toml (Optional) -upgrade (Optional) -h (Optional)")
	flag.PrintDefaults()
}

func printBanner() {
	fmt.Println("")
	banner := "o-o                        \n|  \\                       \n|   O o-o  o-o o-o o-o o-o \n|  /  | | |     \\  |-' |   \no-o   o-o  o-o o-o o-o o   \n                           \n                           "
	fmt.Println(banner)
	fmt.Println("by 0xFTW")
	fmt.Println("")
	fmt.Println("")
}
