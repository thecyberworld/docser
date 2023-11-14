package upgrade

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type Release struct {
	Message    string `json:"message"`
	Name       string `json:"name"`
	TagName    string `json:"tag_name"`
	TarballURL string `json:"tarball_url"`
}

// Upgrade checks the latest version of the docser from the GitHub releases,
// compares it with the current running version, and performs an upgrade if a
// newer version is available. The function uses the GitHub API to fetch the latest
// release information for a given repository. If the tag of the latest release
// does not match the currentVersion, it downloads the new release, builds it and
// replaces the current binary with the new one.
//
// Parameters:
// - currentVersion: A string representing the current version of the application.
// - owner: A string representing the owner or organization of the GitHub repository.
// - repo: A string representing the name of the GitHub repository.
//
// The function does not return any value, but it will print error messages to the
// console if it encounters any issues during the upgrade process, such as errors
// in checking for updates, parsing release data, or upgrading the application.
func Start(currentVersion, owner, repo string) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	// Create a temporary directory to work in (Cross-platform)
	tempDir := os.TempDir()

	// The binary will be stored in the tempDir with the name "docser"
	binaryPath := filepath.Join(tempDir, "docser/docser")

	// Replace the old binary with the new one
	currentBinaryPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to find current binary/executable path: %v", err)
	}

	// Get the latest release
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error checking for updates:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unable to find the latest release, exiting.")
		return
	}

	// Parse the release JSON and store it in the Release struct
	var r Release
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		fmt.Println("Error decoding the release JSON", err)
		return
	}

	// Compare the latest release tag with the current version
	if r.TagName == currentVersion {
		fmt.Println("You are using the latest version.")
		return
	}

	// Download the tarball, extract it, and build the binary
	err = downloadAndBuild(r.TarballURL, tempDir)
	if err != nil {
		fmt.Println("Error when downloading and building:", err)
		return
	}

	// Copy the new binary to the current binary path
	err = copyFile(binaryPath, currentBinaryPath)
	if err != nil {
		fmt.Println("Error copying binary:", err)
		return
	}

	fmt.Println("Application upgraded successfully.")
}

// downloadAndBuild downloads a tarball from the given URL, extracts it, and builds a binary
// from the source code contained in the tarball. It uses a temporary directory specified by
// the tempDir parameter to store the downloaded tarball and the extracted files.
//
// Parameters:
//   - url: The URL from which to download the tarball. It should point directly to the tarball file.
//   - tempDir: The directory where the tarball will be downloaded and extracted. It should have
//     write permissions to allow file creation and modification.
//
// The function returns an error if any step of the process fails, such as downloading the tarball,
// extracting it, or building the binary. Errors are also returned if the glob pattern fails to evaluate
// or if no directories match the pattern for building the binary.
func downloadAndBuild(url, tempDir string) error {
	tarballPath := filepath.Join(tempDir, "docser.tar.gz")

	// Download the tarball
	err := downloadTarball(url, tarballPath)
	if err != nil {
		return fmt.Errorf("error whilst downloading tarball: %w", err)
	}

	// Extract the tarball
	extractedPath := filepath.Join(tempDir, "docser")
	err = extractTarball(tarballPath, extractedPath)
	if err != nil {
		return fmt.Errorf("unable to extract tarball: %w", err)
	}

	pattern := filepath.Join(tempDir, "docser", "*-docser-*")
	directories, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("unable to evaluate glob pattern: %s", err)
	}

	if len(directories) == 0 {
		return fmt.Errorf("no directories found matching pattern: %s", pattern)
	}

	binaryPath := filepath.Join(tempDir, "docser")

	// Build the binary
	cmd := exec.Command("go", "build", "-o", binaryPath)
	cmd.Dir = directories[0]
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("there was a problem when building the updated binary: %w", err)
	}

	return nil
}

// downloadTarball downloads a tarball from the specified URL and saves it to a destination file.
// The function sends a GET request to the provided URL expecting to receive a tarball, which is then
// written to the file system at the specified destination path.
//
// Parameters:
//   - url: The URL where the tarball is located. The URL should be accessible, and the resource should
//     be a valid tarball.
//   - dest: The file path where the downloaded tarball will be saved, including the file name. The
//     directory should be writable.
//
// Returns:
// - An error if there was an issue accessing the URL, creating the destination file, or writing to it.
func downloadTarball(url string, dest string) error {
	// Download the tarball
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("unable to download tarball: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unable to download tarball: %s", url)
	}

	// Create the destination file
	file, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("unable to create file: %w", err)
	}
	defer file.Close()

	// Copy the response body to the destination file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("unable to copy file: %w", err)
	}

	return nil
}

// extractTarball extracts the contents of a gzipped tarball located at the specified file path
// and places the extracted files and directories into the destination directory. The tarball is
// read, and each contained file and directory is created and written to the corresponding location
// within the destination directory, respecting the original file modes and permissions.
//
// Parameters:
//   - tarball: A string representing the file path to the gzipped tarball to be extracted. The tarball
//     must exist and be readable.
//   - dest: A string representing the directory where the tarball's contents will be extracted. The
//     destination directory must be writable and existing files with the same names will be
//     overwritten.
//
// Returns:
//   - nil if the extraction is successful.
//   - An error if there is an issue opening the tarball, reading its contents, creating a gzip reader,
//     or writing extracted files and directories to the destination directory.
func extractTarball(tarball string, dest string) error {
	// Open the tarball file
	file, err := os.Open(tarball)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a gzip reader
	gzr, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("unable to create gzip reader: %w", err)
	}
	defer gzr.Close()

	// Create a tar reader and iterate over the tarball contents
	// Extract each file and directory to the destination directory
	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("unable to read tar.gz file: %w", err)
		}

		target := filepath.Join(dest, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("unable to create directory to extract into: %w", err)
			}
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			f.Close()
		}
	}
	return nil
}

// copyFile copies the content of a source file specified by the `src` parameter to a destination file
// specified by the `dst` parameter. It ensures that the content is successfully written and synced to the
// destination file before returning. If the destination file does not exist, it will be created; if it
// already exists, it will be overwritten.
//
// Parameters:
//   - src: A string representing the path to the source file to be copied. The source file must exist, and
//     be readable.
//   - dst: A string representing the path to the destination file. The destination directory must exist
//     and be writable. If the destination file exists, it will be overwritten.
//
// Returns:
//   - If the function succeeds, it returns nil.
//   - If it fails at any step, such as opening the source file, creating the destination file, copying
//     the content, or syncing the written content to storage, it returns an error detailing what went wrong.
func copyFile(src, dst string) error {
	// Open the source file
	input, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("unable to open source file: %w", err)
	}
	defer input.Close()

	// Create the destination file
	output, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("unable to create destination file: %w", err)
	}
	defer output.Close()

	// Copy the source file to the destination file
	_, err = io.Copy(output, input)
	if err != nil {
		return fmt.Errorf("unable to copy file: %w", err)
	}

	// Sync the destination file
	err = output.Sync()
	if err != nil {
		return fmt.Errorf("unable to sync file: %w", err)
	}

	return nil
}
