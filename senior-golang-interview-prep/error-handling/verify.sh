#!/usr/bin/env bash
set -euo pipefail

here="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
lint="${1:-}"
failures=0
count=0
for qdir in "$here"/[0-9]*-*; do
  [[ -d "$qdir/solution" ]] || continue
  count=$((count + 1))
  echo "== $(basename "$qdir") =="
  if ! (cd "$qdir/solution" && go test ./... && go vet ./...); then
    failures=$((failures + 1))
  fi
  if [[ "$lint" == "--lint" ]]; then
    if command -v staticcheck >/dev/null 2>&1; then
      (cd "$qdir/solution" && staticcheck ./...) || failures=$((failures + 1))
    else
      echo "staticcheck not installed; skipping"
    fi
    if command -v errcheck >/dev/null 2>&1; then
      (cd "$qdir/solution" && errcheck ./...) || failures=$((failures + 1))
    else
      echo "errcheck not installed; skipping"
    fi
  fi
done
echo "Checked $count modules; failures: $failures"
[[ "$count" -eq 60 && "$failures" -eq 0 ]]

