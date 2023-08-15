package scanner

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

func isGitRepository(repositoryPath string) bool {
	repo, err := git.PlainOpen(repositoryPath)
	if err != nil {
		fmt.Printf("[!] Error opening repository: %v\n", err)
		return false
	}

	worktree, err := repo.Worktree()
	if err != nil {
		fmt.Printf("[!] Error getting worktree: %v\n", err)
		return false
	}

	_, err = os.Stat(worktree.Filesystem.Root())
	if err != nil {
		fmt.Printf("[!] Error checking repository path: %v\n", err)
		return false
	}
	return true
}

func InitiateScan(repositoryPath string) {
	if repositoryPath == "." {
		if isGitRepository(repositoryPath) {
			fmt.Printf("[+] Initiating Scan in current directory.")
		}
	} else {
		if isGitRepository(repositoryPath) {
			fmt.Printf("[+] Initiating Scan in %s", repositoryPath)
		}
	}
}
