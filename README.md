# About

Using this project to learn Golang. Until now I have only really known Javascript so this is an opportunity to expand and see what other languages have to offer. Go seems to handle a lot of the shortcomings of Javascript and has a breadth of topics that aren't covered in Javascript.

- Fast and lightweight
- Strongly typed
- Concurrency / Parallelism
- Compiles into a binary
- Pointers

# Goals

1. Create a web server that can authorize a user
2. Enable file upload of pgns that can use the Lichess API to do basic analysis
3. Create a basic frontend to list/sort users games
4. Allow automatic pgn retrieval from chess.com
5. Create a standalone client that can run deeper analysis locally (maybe a separate project)
6. As a separate project, dockerize and run as a GRPC microservice

# Learning notes

## Syntax

- Capital Functions are exported
- For fields to be accessible in structs they also need to be capitalized
- Lowercase functions can still be used by files in the same package
- For a struct to be serialized into json its fields need to tagged `json:field_name` and Marshaled
- The `db:"field"` tag is used to reference table names

## Context

- Context is useful as a per-request place to store information
- Context needs to be keyed by a unique type

## Packages

`go get github.com/...` seems to be the standard to fetch external packages, need to run `go mod vendor` afterwards
`go install ...` the new version? `https://go.dev/doc/go-get-install-deprecation`

## Docs/Notes

- https://pkg.go.dev/github.com/go-chi/chi/v5
- https://pkg.go.dev/github.com/jackc/pgx/v5
- https://go.dev/doc/effective_go
- https://pkg.go.dev/context
- https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
- https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
- https://github.com/benbjohnson/wtf
