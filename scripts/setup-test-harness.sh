#!/usr/bin/env bash

PROJECT_DIR=test-project

mkdir $PROJECT_DIR
cd $PROJECT_DIR
git init
echo "[section]
var1 = 'secret' # nocommit
var2 = 1234
" > setup.conf
git add setup.conf