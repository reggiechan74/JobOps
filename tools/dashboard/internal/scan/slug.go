package scan

import (
	"os"
	"strings"
)

// SplitSlug separates a "{Company}_{Role}_{YYYYMMDD}" stem into a humanized
// title (underscores -> spaces, trailing 8-digit date removed) and the date
// token. The Company/Role boundary is not recoverable, so the whole non-date
// portion is returned as a single title.
func SplitSlug(stem string) (title, date string) {
	parts := strings.Split(stem, "_")
	if n := len(parts); n > 1 && isDate8(parts[n-1]) {
		date = parts[n-1]
		parts = parts[:n-1]
	}
	return strings.Join(parts, " "), date
}

func isDate8(s string) bool {
	if len(s) != 8 {
		return false
	}
	for i := 0; i < 8; i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// DetectCompany returns the name of the Company_Intelligence subfolder whose
// slugified name is a prefix of slug (case-insensitive), or "" if none matches.
func DetectCompany(slug, companyIntelDir string) string {
	entries, err := os.ReadDir(companyIntelDir)
	if err != nil {
		return ""
	}
	lowSlug := strings.ToLower(slug)
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		cslug := strings.ToLower(slugify(e.Name()))
		if strings.HasPrefix(lowSlug, cslug+"_") || lowSlug == cslug {
			return e.Name()
		}
	}
	return ""
}

// slugify converts a company folder name to the underscore form used in slugs.
func slugify(name string) string {
	r := strings.NewReplacer(" ", "_", "&", "_", "-", "_", ",", "", ".", "")
	out := r.Replace(name)
	for strings.Contains(out, "__") {
		out = strings.ReplaceAll(out, "__", "_")
	}
	return strings.Trim(out, "_")
}
