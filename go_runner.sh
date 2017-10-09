#!/usr/bin/env bash

TOOLS="${0%/*}"
BASENAME="${0##*/}"
SCRIPT="${SCRIPT-${BASENAME//git-/}.go}"

go run $TOOLS/$SCRIPT $@
