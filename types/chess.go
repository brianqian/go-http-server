package types

type RawChessComProfile struct {
	Avatar     string `json:"avatar"`
	PlayerId   int    `json:"player_id"`
	Id         string `json:"id"`
	Url        string `json:"url"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Followers  int    `json:"followers"`
	Country    string `json:"country"`
	LastOnline int    `json:"last_online"`
	Joined     int    `json:"joined"`
	Status     string `json:"status"`
	IsStreamer bool   `json:"is_streamer"`
	Verified   bool   `json:"verified"`
	League     string `json:"league"`
}
type ChessComProfile struct {
	Avatar      string `json:"avatar"`
	PlayerId    int    `json:"playerId"`
	Id          string `json:"id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Followers   int    `json:"followers"`
	Country     string `json:"country"`
	Last_online int    `json:"lastOnline"`
	Joined      int    `json:"joined"`
	Status      string `json:"status"`
	IsStreamer  bool   `json:"isStreamer"`
	Verified    bool   `json:"verified"`
	League      string `json:"league"`
}

type PrincipleVariation struct {
	Eval int    `json:"cp"`
	Line string `json:"line"`
	Mate int    `json:"mate"`
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
