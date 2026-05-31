# Dashboard reconcile fixture

`build-fixture.sh` generates `workspace/` — a JobOps workspace with four live
application folders (Acme, Globex, Initech, Umbrella) at different pipeline
stages, company intel for Acme/Umbrella, and a pre-existing `tracker.yaml`
that carries human-status fields (Acme) and a vanished slug (Soylent).

## Verify a reconcile

1. `cd test/fixtures/dashboard && ./build-fixture.sh`
2. Following the algorithm in `plugins/jobops/skills/dashboard/SKILL.md` §Reconcile (created in Task 3),
   produce `workspace/Applications/tracker.yaml`.
3. Normalize the volatile timestamp and compare:
   ```bash
   sed 's/^generated:.*/generated: <RECONCILE_TS>/' \
     workspace/Applications/tracker.yaml > /tmp/actual_tracker.yaml
   diff -u expected_tracker.yaml /tmp/actual_tracker.yaml
   ```
   Expected: no diff.

This asserts: human-zone preservation (Acme), artifact flags, every next_action
branch, archived flip with notes retained (Soylent), new-folder skeletons
(Globex/Initech/Umbrella), and the resume draft/final distinction (Umbrella).
