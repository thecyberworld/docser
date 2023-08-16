package scan_engine

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// StartScanEngine is an exported function from the ScanEngine package
func StartScanEngine(repo *git.Repository, refs []*plumbing.Reference) {
	// Your scan engine logic here
	repository := *repo
	configFunc := repository.Config

	// Call the repository.Config function to get the configuration and error
	repoConfig, err := configFunc()
	if err != nil {
		fmt.Printf("Error getting repository configuration: %v\n", err)
		return
	}
	branches := repoConfig.Branches
	for branch := range branches {
		fmt.Println("Branch here ", branch)
	}
	for _, ref := range refs {
		reference := ref
		fmt.Println("Reference:", reference)
	}
}
