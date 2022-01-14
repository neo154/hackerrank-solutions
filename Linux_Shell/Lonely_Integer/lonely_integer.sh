#!/usr/bin/env bash

read
array=($(cat))
# Simpler than going through the entries array, just use sort and uniq
echo "${array[@]}" | tr ' ' '\n' | sort | uniq -u
