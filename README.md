# go-xz

A small wrapper around libxz to generate lzma2-compressed data from golang

## Installation

go get github.com/fasterthanlime/go-xz

## Usage

import "github.com/fasterthanlime/go-xz"

Then use it like a normal writer to make compressed files:

```go
w := xz.NewWriter(writer)
```

(remember to w.Close() it so that the stream gets flushed)

Then access a compressed stream like a normal reader:

```go
r := xz.NewReader(reader)
```
