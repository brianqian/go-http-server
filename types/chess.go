package types

import "github.com/jackc/pgx/v5/pgtype"

type ChessComProfile struct {
	Avatar      string `json:"avatar"`
	Player_id   int    `json:"playerId"`
	Id          string `json:"id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Followers   int    `json:"followers"`
	Country     string `json:"country"`
	Last_online int    `json:"lastOnline"`
	Joined      int    `json:"joined"`
	Status      string `json:"status"`
	Is_streamer bool   `json:"isStreamer"`
	Verified    bool   `json:"verified"`
	League      string `json:"league"`
}

type PrincipleVariation struct {
	Eval int    `json:"cp"`
	Line string `json:"line"`
}

type ImportedFenParent struct {
	Fen   string               `json:"fen"`
	Evals []ImportedEvaulation `json:"evals"`
}

type ImportedEvaulation struct {
	Pvs    []PrincipleVariation `json:"pvs"`
	Knodes int                  `json:"knodes"`
	Depth  int                  `json:"depth"`
}

type Db_EvalLine struct {
	Id        pgtype.UUID        `json:"id" db:"id"`
	Fen       string             `json:"fen" db:"fen"`
	CreatedAt pgtype.Timestamptz `json:"createdAt" db:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt" db:"updated_at"`
	Line      string             `json:"line" db:"line"`
	Eval      int                `json:"eval" db:"eval"`
	Knodes    int                `json:"knodes" db:"knodes"`
	Depth     int                `json:"depth" db:"depth"`
}
