package patterns

import (
	"bufio"
	"github.com/BurntSushi/toml"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"regexp"
)

// MatchResult represents the result of a regex match
type MatchResult struct {
	FileName    string
	LineNumber  int
	MatchString string
	Pattern     string
}

// Config defines the structure of the TOML config file
type Config struct {
	Patterns map[string]string `toml:"patterns"`
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

	// Load patterns from regex.go
	allPatterns := append([]DefinePatternInfo{}, RegexPatterns...)

	// Load additional patterns from the config file
	if configFile != "" {
		configPatterns, err := loadPatternsFromConfigFile(configFile)
		if err != nil {
			return nil, err
		}
		allPatterns = append(allPatterns, configPatterns...)
	}

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		for _, patternInfo := range allPatterns {
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

func loadPatternsFromConfigFile(configFile string) ([]DefinePatternInfo, error) {
	// Read and parse the TOML config file
	var config Config
	_, err := toml.DecodeFile(configFile, &config)
	if err != nil {
		return nil, err
	}

	var configPatterns []DefinePatternInfo

	// Compile the regex patterns from the config file
	for _, patternStr := range config.Patterns {
		regex, err := regexp.Compile(patternStr)
		if err != nil {
			return nil, err
		}

		// Create a DefinePatternInfo and add it to the configPatterns slice
		configPattern := DefinePatternInfo{
			Pattern:     regex,
			Description: "Config Pattern", // You can set a default description here
		}

		configPatterns = append(configPatterns, configPattern)
	}

	return configPatterns, nil
}
