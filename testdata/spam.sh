#!/bin/bash

# This script echoes "line" on stdout every $DELAY seconds

# USAGE
# $ ./spam.sh $DELAY

DELAY=$1

if test -z $DELAY; then
	echo missing delay arg
	exit 2
fi

while (true); do
	echo line
	sleep $1
done
