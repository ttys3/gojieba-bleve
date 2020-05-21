# GoJieba Bleve support

[![License](https://img.shields.io/badge/license-MIT-yellow.svg?style=flat)](LICENSE)
[![GoDoc](https://godoc.org/github.com/ttys3/gojieba-bleve?status.svg)](https://godoc.org/github.com/ttys3/gojieba-bleve)
[![Go Report Card](https://goreportcard.com/badge/ttys3/gojieba-bleve)](https://goreportcard.com/report/ttys3/gojieba-bleve)

## Intro

GoJieba [bleve](https://github.com/blevesearch/bleve) support mod

this repo exists because [the original mod](https://github.com/yanyiwu/gojieba) has  
[removed bleve support](https://github.com/yanyiwu/gojieba/commit/b714017c2e6807881e0ed64151c0accdd436bb59)

## Get the mod

```bash
go get github.com/ttys3/gojieba-bleve
```

## Usage

import it like this to register the `gojieba` Tokenizer and Analyzer

```go
_ "github.com/ttys3/gojieba-bleve"
```

please see [bleve_test.go](bleve_test.go)