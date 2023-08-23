package main

import (
	"docser/internal/scanner"
	"flag"
	"fmt"
	"os"
)

func main() {

	pRepoLocation := flag.String("d", "", "Directory to be scanned. (Default is current directory)")

	showHelp := flag.Bool("h", false, "Displays help menu")

	flag.Parse()

	if *showHelp {
		showHelpMenu()
		os.Exit(0)
	}

	repositoryPath := *pRepoLocation
	if repositoryPath != "" {
		scanner.InitiateScanandValidatePath(repositoryPath)
	} else {
		scanner.InitiateScanandValidatePath(".")
	}
}

func showHelpMenu() {
	fmt.Println("Usage: docser -d /path/to/directory ")
	flag.PrintDefaults()
}
