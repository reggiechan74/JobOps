#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")" && pwd)/workspace"
rm -rf "$ROOT"
mkdir -p "$ROOT/.jobops"

cat > "$ROOT/.jobops/config.json" <<'JSON'
{
  "version": "2.0",
  "directories": {
    "resume_source": "./ResumeSourceFolder",
    "job_postings": "./Job_Postings",
    "applications_root": "./Applications",
    "company_intelligence": "./Company_Intelligence",
    "career_analysis": "./Career_Analysis",
    "crisis_management": "./Crisis_Management"
  }
}
JSON

# App A — Acme: assessment+resume(final)+cover+osint done; briefing/prep missing -> next: briefing
A="$ROOT/Applications/Acme_SeniorPM_20260520"
mkdir -p "$A/assessment" "$A/resume" "$A/cover-letter" "$A/interview"
echo x > "$A/assessment/assessment.md"
echo x > "$A/resume/step1_draft.md"
echo x > "$A/resume/step2_provenance.md"
echo x > "$A/resume/step3_final.md"
echo x > "$A/cover-letter/cover_letter.md"

# App B — Globex: assessment+resume(draft only); no cover -> resume_final false -> next: buildresume
B="$ROOT/Applications/Globex_DirectorOps_20260524"
mkdir -p "$B/assessment" "$B/resume"
echo x > "$B/assessment/assessment.md"
echo x > "$B/resume/step1_draft.md"

# App C — Initech: empty pipeline -> next: assessjob
C="$ROOT/Applications/Initech_PM_20260528"
mkdir -p "$C"

# App D — Umbrella: everything done -> next: record-outcome
D="$ROOT/Applications/Umbrella_VPProduct_20260510"
mkdir -p "$D/assessment" "$D/resume" "$D/cover-letter" "$D/interview"
echo x > "$D/assessment/assessment.md"
echo x > "$D/resume/step3_final.md"
echo x > "$D/cover-letter/cover_letter.md"
echo x > "$D/interview/briefing.md"
echo x > "$D/interview/interview_prep_part1.md"

# Company intel for Acme + Umbrella only (Globex/Initech have none)
mkdir -p "$ROOT/Company_Intelligence/Acme" "$ROOT/Company_Intelligence/Umbrella"
echo x > "$ROOT/Company_Intelligence/Acme/summary.md"
echo x > "$ROOT/Company_Intelligence/Umbrella/summary.md"

# Pre-existing tracker with human-status fields to prove the merge preserves them,
# plus a vanished slug (Soylent) that must flip to archived.
cat > "$ROOT/Applications/tracker.yaml" <<'YAML'
version: 1
generated: 2026-05-30T09:00:00Z
applications:
  - slug: Acme_SeniorPM_20260520
    company: Acme
    role: Senior PM
    stage: interviewing
    applied_date: 2026-05-21
    next_deadline: 2026-06-02
    contact: "Jane Doe <jane@acme.com>"
    outcome: null
    notes: "Recruiter screen went well"
    artifacts: {assessment: false, resume_draft: false, resume_final: false, cover_letter: false, osint: false, briefing: false, interview_prep: false}
    next_action: assessjob
  - slug: Soylent_LeadPM_20260415
    company: Soylent
    role: Lead PM
    stage: applied
    applied_date: 2026-04-16
    next_deadline: null
    contact: null
    outcome: null
    notes: "Folder archived offline"
    artifacts: {assessment: true, resume_draft: true, resume_final: true, cover_letter: true, osint: false, briefing: false, interview_prep: false}
    next_action: osint
YAML

echo "fixture built at $ROOT"
