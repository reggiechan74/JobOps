// Package model defines the shared types the dashboard engine and its tab
// adapters exchange: a generic Record, its pipeline Stages, and the SkillSpec
// describing a launchable /jobops skill.
package model

import (
	"strings"
	"time"
)

// StageState is the completion state of a single pipeline stage.
type StageState int

const (
	Missing StageState = iota // no artifact on disk
	Draft                     // intermediate artifact only (e.g. resume step1/2)
	Final                     // primary artifact present
)

// Glyph returns the single-rune indicator used in the table/detail views.
func (s StageState) Glyph() string {
	switch s {
	case Final:
		return "✓"
	case Draft:
		return "◐"
	default:
		return "○"
	}
}

// Stage is one step of the JobOps pipeline for a record.
type Stage struct {
	Name  string // e.g. "Resume"
	Skill string // /jobops skill that produces it, e.g. "buildresume"
	State StageState
}

// Lifecycle is the real-world application status the dashboard owns.
type Lifecycle string

const (
	Interested   Lifecycle = "interested"
	Applied      Lifecycle = "applied"
	Screening    Lifecycle = "screening"
	Interviewing Lifecycle = "interviewing"
	Offer        Lifecycle = "offer"
	Accepted     Lifecycle = "accepted"
	Rejected     Lifecycle = "rejected"
	Withdrawn    Lifecycle = "withdrawn"
)

// LifecycleOrder is the cycle order used by the status setter.
var LifecycleOrder = []Lifecycle{
	Interested, Applied, Screening, Interviewing, Offer, Accepted, Rejected, Withdrawn,
}

// Valid reports whether l is a known lifecycle value.
func (l Lifecycle) Valid() bool {
	for _, v := range LifecycleOrder {
		if v == l {
			return true
		}
	}
	return false
}

// Label returns a human-readable title-cased label.
func (l Lifecycle) Label() string {
	if l == "" {
		return ""
	}
	s := string(l)
	return strings.ToUpper(s[:1]) + s[1:]
}

// DefaultLifecycle is the status assumed when no tracker file exists.
func DefaultLifecycle() Lifecycle { return Interested }

// ArgStyle describes how a skill's command line is composed from a record.
type ArgStyle int

const (
	ArgJD      ArgStyle = iota // "<slug>.md --app=<slug>"
	ArgCompany                 // "<company>"
	ArgNone                    // no positional args
)

// SkillSpec describes one launchable skill in a tab's palette.
type SkillSpec struct {
	Name   string // skill name without the prefix, e.g. "assessjob"
	Label  string // display label in the palette
	Arg    ArgStyle
	Plugin string // command namespace; "" means "jobops" (use "jobops-ic" for IC skills)
}

// Record is the generic unit every tab renders.
type Record struct {
	Title     string    // humanized application title (slug minus date)
	Date      string    // trailing date token from the slug, "" if none
	Slug      string    // application slug / folder name
	Company   string    // detected company (for OSINT), "" if unknown
	Score     *int      // normalized fit %, nil if not assessed
	Lifecycle Lifecycle // zero value treated as Interested
	Stages    []Stage   // ordered pipeline
	NextSkill string    // first incomplete stage's skill, "" if complete
	Started   bool      // true if an Applications/{slug}/ folder exists
	Updated   time.Time // most recent artifact mtime (zero if backlog)
	Paths     []string  // backing files/dirs
	Warn      string    // non-empty if slug/frontmatter was malformed
}

// Adapter is implemented by each tab. Phase 1 ships only the Apps adapter.
type Adapter interface {
	Name() string
	Scan() ([]Record, error)
	Skills() []SkillSpec
}
