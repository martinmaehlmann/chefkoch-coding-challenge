#!/usr/bin/env sh
GO_FILES=$(find . -type f -name "*.go" -not -path "./vendor/*")

for dir in $GO_FILES; do
  # check if file contains one or more interfaces
  if grep -q "type.*interface\s{" "$dir"; then
    # get the path without the file
    path_without_file="$(dirname "$dir")"

    # get the filename without the path and without extension
    filename="$(basename "$dir" .go)"

    # extract the package name
    package="$(basename "$path_without_file")"

    # create the mock directory, if it does not already exist
    mkdir -p "$path_without_file/mock/"

    # run mockgen on all identified files
    echo "mockgen -source=$dir -destination=$path_without_file/mock/$filename\_mock.go -package=$package\_mock"
    mockgen -source="$dir" -destination="$path_without_file"/mock/"$filename"\_mock.go -package="$package"_mock
  fi
done