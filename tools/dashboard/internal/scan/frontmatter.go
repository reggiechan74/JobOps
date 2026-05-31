// Package scan walks JobOps data roots and builds dashboard records from the
// files and YAML frontmatter on disk.
package scan

import (
	"bytes"
	"os"
	"strconv"
	"strings"

	"github.com/adrg/frontmatter"
)

// AppsMatter holds the frontmatter fields the Apps adapter cares about.
type AppsMatter struct {
	OutputType      string `yaml:"output_type"`
	Status          string `yaml:"status"`
	NormalizedScore string `yaml:"normalized_score"`
	OverallScore    string `yaml:"overall_score"`
}

// ParseFrontmatter parses the YAML frontmatter of a markdown file into out.
// A missing or malformed frontmatter block is returned as a non-empty warn
// string rather than a hard error, so a scan never aborts on one bad file.
func ParseFrontmatter(path string, out interface{}) (warn string, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	// Detect whether the file begins with a frontmatter delimiter.
	// adrg/frontmatter silently returns nil on a body with no "---" block,
	// so we must guard this ourselves to surface a useful warning.
	if !strings.HasPrefix(strings.TrimLeft(string(data), "\r\n"), "---") {
		return "no frontmatter block", nil
	}

	if _, perr := frontmatter.Parse(bytes.NewReader(data), out); perr != nil {
		return "could not parse frontmatter: " + perr.Error(), nil
	}
	return "", nil
}

// leadingInt extracts the leading run of digits from s (e.g. "84%" -> 84,
// "168/200" -> 168). Returns ok=false when s has no leading digit.
func leadingInt(s string) (int, bool) {
	end := 0
	for end < len(s) && s[end] >= '0' && s[end] <= '9' {
		end++
	}
	if end == 0 {
		return 0, false
	}
	n, err := strconv.Atoi(s[:end])
	if err != nil {
		return 0, false
	}
	return n, true
}
