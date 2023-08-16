package scan_engine

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
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
	fmt.Print(repoConfig.Branches)
	// Open the repository's HEAD reference to get the reference's hash
	headRef, err := repo.Head()
	if err != nil {
		fmt.Printf("Error getting HEAD reference: %v\n", err)
		return
	}

	// Retrieve the hash of the HEAD reference
	headHash := headRef.Hash()

	// Get the commit object of the HEAD reference
	commit, err := repo.CommitObject(headHash)
	if err != nil {
		fmt.Printf("Error getting commit: %v\n", err)
		return
	}

	// Iterate through each commit in the repository
	commitIter, err := repo.Log(&git.LogOptions{From: commit.Hash})
	if err != nil {
		fmt.Printf("Error getting commit history: %v\n", err)
		return
	}
	defer commitIter.Close()

	err = commitIter.ForEach(func(commitObj *object.Commit) error {
		fmt.Println("Commit:", commitObj.Hash)

		// Access the files in the commit using commitObj.Files() and iterate through them
		fileIter, err := commitObj.Files()
		if err != nil {
			fmt.Printf("Error getting commit files: %v\n", err)
			return err
		}
		defer fileIter.Close()

		err = fileIter.ForEach(func(file *object.File) error {
			fmt.Println("File:", file.Name)
			return nil
		})
		if err != nil {
			fmt.Printf("Error iterating through files: %v\n", err)
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Error iterating commits: %v\n", err)
		return
	}
	for _, ref := range refs {
		reference := ref
		fmt.Println("Reference:", reference)
	}
}
