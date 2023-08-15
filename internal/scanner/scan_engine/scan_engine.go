package scan_engine

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// StartScanEngine is an exported function from the ScanEngine package
func StartScanEngine(repo *git.Repository, refs []*plumbing.Reference) {
	// Your scan engine logic here
	fmt.Println("StartScanEngine called", repo, refs)
}
