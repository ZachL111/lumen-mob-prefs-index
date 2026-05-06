# Lumen Mob Prefs Index Walkthrough

I use this file as a small checklist before changing the Go implementation.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | form pressure | 212 | ship |
| stress | sync drift | 178 | ship |
| edge | local state | 216 | ship |
| recovery | conflict cost | 211 | ship |
| stale | form pressure | 191 | ship |

Start with `edge` and `stress`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

The useful comparison is `local state` against `sync drift`, not the raw score alone.
