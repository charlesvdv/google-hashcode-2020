#!/usr/bin/env bash

set -e

mkdir -p output

for input in ./input/*.in; do
    filename=$(basename "$input" .in)
    echo "Processing $filename..."
    go run main.go < $(echo $input) > "./output/$filename.out"
done