#!/usr/bin/env python3
"""Verify runnable Go examples and local Markdown links in the LeetCode learning library.

Run from any directory:
    python3 leetcode/0-problem-solving-building-blocks/verify-examples.py

Requirements: Python 3.9+, Go in PATH. The program writes Go code only to a
temporary directory and does not edit the repository.
"""
from __future__ import annotations

import re
import shutil
import subprocess
import sys
import tempfile
from pathlib import Path
from urllib.parse import unquote, urlsplit


def main() -> int:
    foundations = Path(__file__).resolve().parent
    leetcode = foundations.parent
    lessons = sorted(foundations.glob("[0-9][0-9]-*/README.md"))
    expected_count = 20
    if len(lessons) != expected_count:
        print(f"ERROR: expected {expected_count} lessons; found {len(lessons)}")
        return 1
    if shutil.which("go") is None:
        print("ERROR: Go compiler missing from PATH")
        return 1

    failures: list[str] = []
    successful = 0

    for path in [foundations / "README.md", leetcode / "README.md", *lessons]:
        contents = path.read_text(encoding="utf-8")
        for raw_link in re.findall(r"\]\(([^)]+)\)", contents):
            parsed = urlsplit(raw_link)
            if parsed.scheme or raw_link.startswith("#"):
                continue
            target = unquote(parsed.path)
            if not target:
                continue
            if not (path.parent / target).resolve().is_file():
                failures.append(f"Broken Markdown link: {path.relative_to(leetcode)} → {raw_link}")

    for readme in lessons:
        contents = readme.read_text(encoding="utf-8")
        examples = re.findall(r"^" + chr(96) * 3 + r"go\s*\n(.*?)^" + chr(96) * 3, contents, re.M | re.S)
        label = readme.parent.name
        if len(examples) != 1:
            failures.append(f"{label}: expected exactly one Go example, found {len(examples)}")
            continue
        snippet = examples[0]
        if not snippet.lstrip().startswith("package main"):
            failures.append(f"{label}: snippet must be complete (package main)")
            continue
        with tempfile.TemporaryDirectory(prefix="go-foundation-") as tmp:
            source = Path(tmp) / "main.go"
            source.write_text(snippet + "\n", encoding="utf-8")
            try:
                result = subprocess.run(
                    ["go", "run", source.name],
                    cwd=tmp,
                    text=True,
                    capture_output=True,
                    timeout=90,
                    check=False,
                    env=__import__("os").environ | {"GO111MODULE": "off"},
                )
            except subprocess.TimeoutExpired:
                failures.append(f"{label}: example timed out")
                continue
            if result.returncode:
                failures.append(f"{label}: Go failed:\n{result.stderr}")
            else:
                successful += 1
                print(f"PASS {label}: {result.stdout.strip().splitlines()[-1:] or ['(no output)']}")

    if failures:
        print(f"\n{len(failures)} verification problem(s):")
        for problem in failures:
            print(f"- {problem}")
        return 1

    print(f"\nAll {successful}/{expected_count} foundation Go examples ran; local Markdown targets exist.")
    print("This checks compilation/execution and link paths, not correctness of the LeetCode solutions.")
    return 0


if __name__ == "__main__":
    sys.exit(main())
