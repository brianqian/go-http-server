package types

import "net/http"

type key int

const UserIdKey key = iota + 1

const (
	ChessComAPI string = "https://api.chess.com/"
	LichessAPI  string = "https://lichess.org/api/"
)

type RequestOpts struct {
	Body    interface{}
	Headers map[string][]string
}

type HttpClient interface {
	Get(url string) ([]byte, *http.Response, error)
	Post(url string, body interface{}) ([]byte, *http.Response, error)
	Put(url string, body interface{}) ([]byte, *http.Response, error)
	Delete(url string, body interface{}) ([]byte, *http.Response, error)
	MakeRequest(method string, url string, opts RequestOpts) ([]byte, *http.Response, error)
}
