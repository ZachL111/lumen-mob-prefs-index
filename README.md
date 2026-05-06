# lumen-mob-prefs-index

`lumen-mob-prefs-index` keeps a focused Go implementation around mobile workflows. The project goal is to create a Go reference implementation for prefs workflows, centered on state machine modeling, transition tables, and invalid-transition tests.

## Why I Keep It Small

This is intentionally local and self-contained so it can be inspected without credentials, services, or seeded history.

## Lumen Mob Prefs Index Review Notes

For a quick review, compare `local state` with `sync drift` before reading the middle cases.

## Included Behavior

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/lumen-mob-prefs-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `local state` and `sync drift`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Internal Model

The repository has two validation layers: the original compact policy fixture and the domain review fixture. They are separate so one can change without hiding failures in the other.

The Go addition stays small enough to inspect in one sitting.

## Try It Locally

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Validation

The verifier is intentionally local. It should fail if the fixture score math, lane assignment, or language-specific test drifts.

## Scope

No external service is required. A deeper version would add more negative cases and a clearer boundary around invalid input.
