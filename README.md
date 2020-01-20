# go-fileparallelread

This go module can be used to read a single file in parallel.

## How it works

- Given a file, multiple section readers are created to read it in parallel. The section readers may or may not use multiple file handles at the os level; this is configurable.
- The section readers are created using a delimitter provided by the user

## Usage
Checkout the API documentation at https://godoc.org/

## License
Copyright 2020 Harleen Mann. All rights reserved.<br />
Use of this source code is governed by a GNU GENERAL PUBLIC LICENSE that can be found in the LICENSE file