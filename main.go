package main

import (
	"docser/internal/scanner"
	"flag"
	"fmt"
	"os"
)

func main() {

	pRepoLocation := flag.String("d", "", "Help message for directory argument")

	showHelp := flag.Bool("h", false, "Help Menu")

	flag.Parse()

	if *showHelp {
		showHelpMenu()
		os.Exit(0)
	}

	repositoryPath := *pRepoLocation
	if repositoryPath != "" {
		scanner.InitiateScanandValidatePath(repositoryPath)
	} else {
		showHelpMenu()
	}
}

func showHelpMenu() {
	fmt.Println("Usage: docser [options] subcommand")
	flag.PrintDefaults()
}
