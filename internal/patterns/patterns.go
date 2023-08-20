package patterns

import (
	"io"
	"strings"

	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// MatchResult represents the result of a regex match
type MatchResult struct {
	FileName    string
	LineNumber  int
	MatchString string
	Pattern     string
}

// ProcessTextFileContentsWithRegex reads and processes the contents of a text-based file using regex patterns
func ProcessTextFileContentsWithRegex(file *object.File) ([]MatchResult, error) {
	fileReader, err := file.Reader()
	if err != nil {
		return nil, err
	}
	defer fileReader.Close()

	fileContents, err := io.ReadAll(fileReader)
	if err != nil {
		return nil, err
	}

	var matchResults []MatchResult

	// Iterate through each line and check for regex matches
	lines := strings.Split(string(fileContents), "\n")
	for lineNumber, line := range lines {
		for _, patternInfo := range RegexPatterns {
			regex := patternInfo.Pattern
			if regex.MatchString(line) {
				submatches := regex.FindStringSubmatch(line)
				if len(submatches) > 0 {
					matchResult := MatchResult{
						FileName:    file.Name,
						LineNumber:  lineNumber + 1, // Line numbers are 1-based
						MatchString: submatches[0],
						Pattern:     patternInfo.Description,
					}
					matchResults = append(matchResults, matchResult)
				}
			}
		}
	}

	return matchResults, nil
}
