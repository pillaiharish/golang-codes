#!/usr/bin/env python3
"""Validate the practice library and run starters; does not grade solutions."""
import concurrent.futures
import json
import os
from pathlib import Path
import re
import shutil
import subprocess
import sys
from urllib.parse import unquote

BASE = Path(__file__).resolve().parent
manifest = json.loads((BASE / 'manifest.json').read_text())
problems = manifest['problems']
assert len(problems) == 150, 'Expected the NeetCode 150 membership snapshot'
assert len({p['id'] for p in problems}) == 150, 'Duplicate problem IDs'
assert len({p['folder'] for p in problems}) == 150, 'Duplicate folders'
actual = {p.name for p in BASE.iterdir() if p.is_dir() and re.match(r'^\d+\. ', p.name)}
assert actual == {p['folder'] for p in problems}, 'Folder/manifest mismatch'
concepts = sorted((BASE / '0-concepts').glob('*/main.go'))
assert len(concepts) == 24, 'Expected 24 concept demonstrations'
for problem in problems:
    directory = BASE / problem['folder']
    assert (directory / 'main.go').exists()
    assert (directory / 'README.md').exists()
    assert problem['concepts'], f"Missing concepts: {problem['id']}"
    for key in problem['concepts']:
        lesson = BASE / '0-concepts' / key / 'README.md'
        assert lesson.exists(), key
        assert f"#{problem['id']} " in lesson.read_text(), 'Missing concept backlink'

link_count = 0
for doc in BASE.rglob('*.md'):
    for link in re.findall(r'\]\(([^)]+)\)', doc.read_text()):
        if re.match(r'\w+://', link):
            continue
        path, _, anchor = unquote(link).partition('#')
        target = (doc.parent / path).resolve() if path else doc
        assert target.exists(), f'{doc}: missing {link}'
        if anchor and target.suffix == '.md':
            headings = re.findall(r'^#+ (.+)$', target.read_text(), re.M)
            anchors = {re.sub(r'[^\w\- ]', '', h.lower()).replace(' ', '-') for h in headings}
            assert anchor in anchors, f'{doc}: missing section {link}'
        link_count += 1

go = shutil.which('go')
if not go:
    sys.exit('Navigation checks passed, but go is missing from PATH; no execution claim can be made.')
print(f'150 folders, 24 concepts, {link_count} local links checked.', flush=True)
print(subprocess.check_output([go, 'version'], text=True).strip(), flush=True)
env = dict(os.environ, GOTOOLCHAIN='local')

def run(source):
    try:
        result = subprocess.run([go, 'run', 'main.go'], cwd=source.parent,
                                env=env, capture_output=True, text=True, timeout=90)
        if result.returncode:
            return str(source.parent.relative_to(BASE)), result.stderr, False
        return str(source.parent.relative_to(BASE)), result.stdout, True
    except subprocess.TimeoutExpired:
        return str(source.parent.relative_to(BASE)), 'Execution exceeded 90 seconds', False

sources = [BASE / p['folder'] / 'main.go' for p in problems] + concepts
failures = []
todos = 0
with concurrent.futures.ThreadPoolExecutor(max_workers=min(4, os.cpu_count() or 1)) as pool:
    for i, (name, output, ok) in enumerate(pool.map(run, sources), 1):
        if not ok:
            failures.append((name, output))
        elif 'TODO: implement' in output:
            todos += 1
        if i % 25 == 0:
            print(f'Executed {i}/{len(sources)} entry points', flush=True)
for name, output in failures:
    print(f'FAILED {name}\n{output}', file=sys.stderr)
if failures:
    sys.exit(1)
print(f'All {len(sources)} entry points ran; {todos} reported TODO (unsolved).')
print('This checks scaffolding/execution only, not algorithm correctness or learner progress.')
