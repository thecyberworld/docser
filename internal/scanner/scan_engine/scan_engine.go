package scan_engine

import (
	"fmt"
	"io"
	"log"

	"docser/internal/patterns"

	"gopkg.in/h2non/filetype.v1"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// StartScanEngine is an exported function from the ScanEngine package
func StartScanEngine(repo *git.Repository, refs []*plumbing.Reference) {
	repository := *repo
	configFunc := repository.Config

	// Call the repository.Config function to get the configuration and error
	_, err := configFunc()
	if err != nil {
		log.Printf("Error getting repository configuration: %v\n", err)
		return
	}

	// Open the repository's HEAD reference to get the reference's hash
	headRef, err := repo.Head()
	if err != nil {
		log.Printf("Error getting HEAD reference: %v\n", err)
		return
	}

	// Retrieve the hash of the HEAD reference
	headHash := headRef.Hash()

	// Get the commit object of the HEAD reference
	commit, err := repo.CommitObject(headHash)
	if err != nil {
		log.Printf("Error getting commit: %v\n", err)
		return
	}

	// Iterate through each commit in the repository
	err = iterateCommits(repo, commit, refs)
	if err != nil {
		log.Printf("Error iterating commits: %v\n", err)
		return
	}
}

// iterateCommits iterates through each commit and processes files
func iterateCommits(repo *git.Repository, commit *object.Commit, refs []*plumbing.Reference) error {
	commitIter, err := repo.Log(&git.LogOptions{From: commit.Hash})
	if err != nil {
		return err
	}
	defer commitIter.Close()

	return commitIter.ForEach(func(commitObj *object.Commit) error {

		_ = getBranchName(commitObj.Hash, refs)

		// Access and process the files in the commit
		err := processCommitFiles(commitObj)
		if err != nil {
			return err
		}

		return nil
	})
}

// getBranchName returns the name of the branch that the commit belongs to
func getBranchName(commitHash plumbing.Hash, refs []*plumbing.Reference) string {
	foundBranch := "Unknown"
	for _, ref := range refs {
		if ref.Hash() == commitHash {
			foundBranch = ref.Name().Short()
			break
		}
	}
	return foundBranch
}

// processCommitFiles accesses and processes the files in the commit
func processCommitFiles(commitObj *object.Commit) error {
	fileIter, err := commitObj.Files()
	if err != nil {
		return err
	}
	defer fileIter.Close()

	return fileIter.ForEach(func(file *object.File) error {

		result, err := patterns.ProcessTextFileContentsWithRegex(file)
		if (err == nil) && (len(result) != 0) {
			fmt.Println(commitObj.Hash)
			fmt.Println("File:", file.Name)
		}

		// Check if the file extension corresponds to text-based formats
		if isTextFile(file) {
			// Access and process the contents of the file
			err := processTextFileContents(file)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// processTextFileContents call file processing and regex matching function
func processTextFileContents(file *object.File) error {
	result, err := patterns.ProcessTextFileContentsWithRegex(file)
	if (err == nil) && (len(result) != 0) {
		log.Println("Result", result)
	}
	return nil
}

// isTextFile checks if the file extension corresponds to a text-based format by checking magic byte of the file
func isTextFile(file *object.File) bool {
	fileReader, err := file.Reader() // Assuming 'Reader' is a method in your 'object.File' type
	if err != nil {
		log.Printf("Error opening %s: %v\n", file.Name, err)
		return false
	}
	defer func(fileReader io.ReadCloser) {
		err := fileReader.Close()
		if err != nil {
			log.Fatalln("[!] Error deferring File Reader")
		}
	}(fileReader)

	bufferSize := 261
	buffer := make([]byte, bufferSize) // Read the first 261 bytes for magic number detection
	bufLen, err := fileReader.Read(buffer)
	if (err != nil) && (bufLen > bufferSize) {
		log.Printf("Error reading %s: %v\n", file.Name, err)
		return false
	}

	kind, _ := filetype.Match(buffer)

	return kind == filetype.Unknown
}
