package main

import (
	"docser/internal/scanner"
	"flag"
	"fmt"
	"os"
)

func main() {

	pRepoLocation := flag.String("d", "", "Directory to be scanned. (Default is current directory)")
	pConfigFile := flag.String("c", "", "Docser config file (Must end in .toml)")
	showHelp := flag.Bool("h", false, "Displays help menu")

	flag.Parse()

	if *showHelp {
		showHelpMenu()
		os.Exit(0)
	}
	printBanner()
	repositoryPath := *pRepoLocation
	configFile := *pConfigFile
	scanner.ParseConfigAndInitiateScan(configFile, repositoryPath)
}

func showHelpMenu() {
	fmt.Println("Usage: docser -d /path/to/directory -c /path/to/.docser.toml (Optional)")
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
