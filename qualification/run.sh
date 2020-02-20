#!/usr/bin/env bash

set -e

mkdir -p output

executable="$(basename $(pwd)).out"
echo -n "Compiling... "
go build -o $executable
echo "done"

for input in ./input/*; do
    filename=$(basename "$input" .in)
    echo -n "Processing $filename... "
    bash -c "./$executable" < $(echo $input) > "./output/$filename.out"
    echo "done"
done

echo "Zipping code... "
zip output/code.zip *.go