package types

type key int

const UserIdKey key = iota + 1

const (
	ChessComAPI string = "https://api.chess.com/"
	LichessAPI  string = "https://lichess.org/api/"
)

type HttpClient interface {
	Get(url string) ([]byte, error)
	Post(url string, body interface{}) ([]byte, error)
	Put(url string, body interface{}) ([]byte, error)
	Delete(url string, body interface{}) (bool, error)
}
