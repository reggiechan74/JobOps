package scan

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSplitSlug(t *testing.T) {
	cases := []struct {
		in          string
		title, date string
	}{
		{"Acme_Product_Manager_20260531", "Acme Product Manager", "20260531"},
		{"Beta_Capital_Analyst", "Beta Capital Analyst", ""},
		{"Gamma_20251231", "Gamma", "20251231"},
	}
	for _, c := range cases {
		title, date := SplitSlug(c.in)
		if title != c.title || date != c.date {
			t.Errorf("SplitSlug(%q) = (%q,%q), want (%q,%q)", c.in, title, date, c.title, c.date)
		}
	}
}

func TestDetectCompany(t *testing.T) {
	root := t.TempDir()
	intel := filepath.Join(root, "Company_Intelligence", "Acme")
	if err := os.MkdirAll(intel, 0o755); err != nil {
		t.Fatal(err)
	}
	got := DetectCompany("Acme_Product_Manager_20260531", filepath.Join(root, "Company_Intelligence"))
	if got != "Acme" {
		t.Errorf("DetectCompany = %q, want Acme", got)
	}
	if DetectCompany("Zzz_Role_20260101", filepath.Join(root, "Company_Intelligence")) != "" {
		t.Errorf("expected empty company for non-matching slug")
	}
}
