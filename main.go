package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

func main() {

	pRepoLocation := flag.String("d", ".", "Help message for directory argument")

	showHelp := flag.Bool("h", false, "Help Menu")

	flag.Parse()

	if *showHelp {
		showHelpMenu()
		os.Exit(0)
	}

	repositoryPath := *pRepoLocation

	if repositoryPath == "." {
		isGitRepository(repositoryPath)
		fmt.Printf("[+] Initiating Scan in current directory.")
	} else {
		fmt.Printf("[+] Initiating Scan in %s", repositoryPath)
	}
}

func showHelpMenu() {
	fmt.Println("Usage: docser [options] subcommand")
	flag.PrintDefaults()
}

func isGitRepository(repositoryPath string) bool {
	repo, err := git.PlainOpen(repositoryPath)
	if err != nil {
		fmt.Printf("Error opening repository: %v\n", err)
		return false
	}

	worktree, err := repo.Worktree()
	if err != nil {
		fmt.Printf("Error getting worktree: %v\n", err)
		return false
	}

	_, err = os.Stat(worktree.Filesystem.Root())
	if err != nil {
		fmt.Printf("Error checking repository path: %v\n", err)
		return false
	}
	return true
}
