package scanner

import (
	"docser/internal/scanner/scan_engine"
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
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

func InitiateScanandValidatePath(repositoryPath string) {
	if repositoryPath == "." {
		if isGitRepository(repositoryPath) {
			fmt.Printf("[+] Initiating Scan in current directory.\n")
			startScanEngine(repositoryPath)
		}
	} else {
		if isGitRepository(repositoryPath) {
			fmt.Printf("[+] Initiating Scan in %s \n", repositoryPath)
			startScanEngine(repositoryPath)
		}
	}
}

func startScanEngine(repositoryPath string) {
	repo, err := git.PlainOpen(repositoryPath)
	if err != nil {
		fmt.Printf("[!] Error opening repository: %v\n", err)
		return
	}

	refs, err := repo.References()
	if err != nil {
		fmt.Printf("[!] Error getting references: %v\n", err)
		return
	}

	var plumbingRefs []*plumbing.Reference
	err = refs.ForEach(func(ref *plumbing.Reference) error {
		plumbingRefs = append(plumbingRefs, ref)
		return nil
	})
	if err != nil {
		fmt.Printf("[!] Error iterating through references: %v\n", err)
		return
	}
	// Call the StartScanEngine function from the ScanEngine package
	scan_engine.StartScanEngine(repo, plumbingRefs)
}
