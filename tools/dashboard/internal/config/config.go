// Package config loads the JobOps workspace configuration (.jobops/config.json)
// and resolves the configured directories to absolute paths.
package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// Config mirrors the parts of .jobops/config.json the dashboard reads.
type Config struct {
	Directories struct {
		ResumeSource        string `json:"resume_source"`
		JobPostings         string `json:"job_postings"`
		ApplicationsRoot    string `json:"applications_root"`
		CompanyIntelligence string `json:"company_intelligence"`
		CareerAnalysis      string `json:"career_analysis"`
		CrisisManagement    string `json:"crisis_management"`
	} `json:"directories"`

	Root string `json:"-"` // workspace root (the dir that contains .jobops/)
}

// Discover walks up from start looking for the nearest ancestor that contains a
// .jobops directory, and returns that ancestor path.
func Discover(start string) (string, error) {
	dir, err := filepath.Abs(start)
	if err != nil {
		return "", err
	}
	for {
		info, err := os.Stat(filepath.Join(dir, ".jobops"))
		if err == nil && info.IsDir() {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("no .jobops directory found in any parent of " + start)
		}
		dir = parent
	}
}

// Load reads .jobops/config.json from the given workspace root.
func Load(root string) (*Config, error) {
	data, err := os.ReadFile(filepath.Join(root, ".jobops", "config.json"))
	if err != nil {
		return nil, err
	}
	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	c.Root = root
	return &c, nil
}

// resolve joins a configured (possibly relative, possibly "./"-prefixed) path
// with the workspace root, unless it is already absolute.
func (c *Config) resolve(p string) string {
	if p == "" {
		return ""
	}
	if filepath.IsAbs(p) {
		return p
	}
	return filepath.Join(c.Root, filepath.Clean(p))
}

func (c *Config) JobPostingsDir() string  { return c.resolve(c.Directories.JobPostings) }
func (c *Config) ApplicationsDir() string { return c.resolve(c.Directories.ApplicationsRoot) }
func (c *Config) CompanyIntelDir() string { return c.resolve(c.Directories.CompanyIntelligence) }
