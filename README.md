# go-kafka

> An abstract middleware that unifies Apache Kafka message processing across various libraries

[![build](https://img.shields.io/github/actions/workflow/status/alebabai/go-kafka/ci.yml)](https://github.com/alebabai/go-kafka/actions?query=workflow%3ACI)
[![version](https://img.shields.io/github/go-mod/go-version/alebabai/go-kafka)](https://go.dev/)
[![report](https://goreportcard.com/badge/github.com/alebabai/go-kafka)](https://goreportcard.com/report/github.com/alebabai/go-kafka)
[![coverage](https://img.shields.io/codecov/c/github/alebabai/go-kafka)](https://codecov.io/github/alebabai/go-kafka)
[![tag](https://img.shields.io/github/tag/alebabai/go-kafka.svg)](https://github.com/alebabai/go-kafka/tags)
[![reference](https://pkg.go.dev/badge/github.com/alebabai/go-kafka.svg)](https://pkg.go.dev/github.com/alebabai/go-kafka)

## Getting started

Go modules are supported.  

Manual install:

```bash
go get -u github.com/alebabai/go-kafka
```

Golang import:

```go
import "github.com/alebabai/go-kafka"
```

## Usage

To use the abstractions provided by this module, please implement the converters defined in [adapter/converter.go](./adapter/converter.go) for the types specific to your Apache Kafka client library.

Additionally, here is an adapter implementation for the most popular Apache Kafka client library, [github.com/Shopify/sarama](github.com/Shopify/sarama). The module, named [github.com/alebabai/go-kafka/adapter/sarama](./adapter/sarama), can be used independently or as a reference example.
