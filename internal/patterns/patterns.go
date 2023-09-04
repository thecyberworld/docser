package patterns

import (
	"bufio"

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
func ProcessTextFileContentsWithRegex(file *object.File, configFile string) ([]MatchResult, error) {
	fileReader, err := file.Reader()
	if err != nil {
		return nil, err
	}
	defer fileReader.Close()

	var matchResults []MatchResult

	scanner := bufio.NewScanner(fileReader)
	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		for _, patternInfo := range RegexPatterns {
			regex := patternInfo.Pattern
			if regex.MatchString(line) {
				submatches := regex.FindStringSubmatch(line)
				if len(submatches) > 0 {
					matchResult := MatchResult{
						FileName:    file.Name,
						LineNumber:  lineNumber, // Line numbers are 1-based
						MatchString: submatches[0],
						Pattern:     patternInfo.Description,
					}
					matchResults = append(matchResults, matchResult)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matchResults, nil
}
