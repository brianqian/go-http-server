# About

Using this project to learn Golang. Until now I have only really known Javascript so this is an opportunity to expand and see what other languages have to offer. Go seems to handle a lot of the shortcomings of Javascript and has a breadth of topics that aren't covered in Javascript.

- Fast and lightweight
- Strongly typed
- Concurrency
- Compiles into a binary
- Pointers

# Choices

## Chi Router vs http

# Learning notes

## Syntax

- Capital Functions are exported
- For a struct to be serialized into json its fields need to labeled `json:field_name` and Marshaled

## Context

- Context is useful as a per-request place to store information
- Context needs to be keyed by a specific type

## Packages

`go get github.com/...` seems to be the standard to fetch external packages, need to run `go mod vendor` afterwards

## Docs/Notes

- https://pkg.go.dev/github.com/go-chi/chi/v5
- https://pkg.go.dev/github.com/jackc/pgx/v5
- https://go.dev/doc/effective_go
- https://pkg.go.dev/context
- https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
- https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
- https://github.com/benbjohnson/wtf