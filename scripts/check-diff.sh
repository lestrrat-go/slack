#!/bin/sh

DIFF=$(git diff)

if [ -z "$DIFF" ]; then
    exit 0
fi

echo "$DIFF"
echo ""
echo "Above diff was found."

exit 1