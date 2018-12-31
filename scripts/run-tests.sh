#!/usr/bin/env bash

cd test-project
TESTS=1
PASSES=0
FAILS=0

# TEST 1
git commit -m "test"
if [ $? -eq 1 ]
then
	((PASSES++))
else
	((FAILS++))
fi

# TEST RESULTS
echo
echo "## ############ ##"
echo "## TEST RESULTS ##"
echo "## ############ ##"
echo "## PASSED: $PASSES, FAILED: $FAILS"

if [[ "$FAILS" -gt 0 ]]
then
	echo "## TEST FAILED"
	echo "## ############ ##"
	exit 1
fi

echo "## ############ ##"
exit 0