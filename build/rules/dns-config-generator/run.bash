#!/usr/bin/env bash
set -e

BIN=@@BIN@@
SRC=@@SRC@@
DIR=@@DIR@@
OUTFILE_NAME=@@OUTFILE@@

RUNFILES=$(pwd)

COMMAND="$RUNFILES/$BIN"
SOURCE="$RUNFILES/$SRC"

cd "$BUILD_WORKSPACE_DIRECTORY"/"$DIR"
"$COMMAND" "$SOURCE" > $OUTFILE_NAME