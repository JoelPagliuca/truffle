# Nocommit
Utility for preventing secrets being committed to repositories

### Notes

```py
# these will give all filenames of changed files
GIT_DIFF = "git diff --cached --diff-filter ACMU --name-only"
GIT_DIFF = "git diff-tree -r --no-commit-id --name-only --diff-filter ACMU {old} {new}"
# this will read in the file being pushed to the server
GIT_SHOW = "git show {ref}:{filename}"
DEFAULT_GIT_REF = "4b825dc642cb6eb9a060e54bf8d69288fbee4904"

```

