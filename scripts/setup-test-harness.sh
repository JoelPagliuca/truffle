#!/usr/bin/env bash

mkdir test-project
cd test-project
git init
echo "[section]
var1 = 'secret' # nocommit
var2 = 1234
" > setup.conf
git add setup.conf
cp ../nocommit ./.git/hooks/pre-commit