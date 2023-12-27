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

## Importing large files

Lichess provides a 4.5 gb json file of pre-calculated chess positions. The process of doing this requires reading the file, unmarshalling each line into a struct, and then ingesting it into the database. This seemed like a good opportunity to use some goroutines and look at memory management

### Unmarshalling

- The first attempt was just to read the file one line at a time and store it in a struct. This took around 45 seconds
- The next idea was to use some goroutines to break this work up across multiple threads but learned that
  - Work like this is much more i/o intensive than CPU bound (all though it seems like single cores are maxed?)
  - Goroutines make the import slower. Because they operate on different threads, the CPU has different caches that need to be updated by each thread, slowing down the whole operation.
- `bufio` has a default buffer of around 4kb per token so memory isn't an issue
- A next optimization could be using `sync.Pool` to recycle data structures

### DB ingestion

- There was a clear difference here using goroutines instead of a sync upload since we could use multiple connections to start importing the data. However halfway through the process RAM gets maxed out, cpu quickly follows and the import slows to a crawl.
- I refactored the database insert to take in multiple values per query and used pgx's batch feature to make each transaction more efficient.
- Since bufio scans files one line at a time, I used `chanx` to batch together each line to feed the bulk insert function
- Once batches exceeded the number of db connections there seems to be a big slowdown. The last steps are to
  - See if the code can be modified to at least wait for the context timeout before ending
  - Use the CopyWith method to use postgres copy to ingest a large file

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
