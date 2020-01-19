# go-fileparallelread

This go module shows how to read a single file in parallel.

## How it works
Given a file, multiple section readers are created to read it in parallel. The section readers may or may not use multiple file handles at the os level; this is configurable.
The section readers are created using a delimitter provided by the user

## Usage
Checkout the API documentation at https://godoc.org/
