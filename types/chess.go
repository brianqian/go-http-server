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

type pgnPlayerData struct {
	Rating   int    `json:"rating"`
	Result   string `json:"result"`
	Id       string `json:"@id"`
	Username string `json:"username"`
	Uuid     string `json:"uuid"`
}
type RawChessComPgn struct {
	Url          string        `json:"url"`
	Pgn          string        `json:"pgn"`
	TimeControl  string        `json:"time_control"`
	EndTime      int           `json:"end_time"`
	Rated        bool          `json:"rated"`
	Tcn          string        `json:"tcn"`
	Uuid         string        `json:"uuid"`
	InitialSetup string        `json:"initial_setup"`
	Fen          string        `json:"fen"`
	TimeClass    string        `json:"time_class"`
	Rules        string        `json:"rules"`
	White        pgnPlayerData `json:"white"`
	Black        pgnPlayerData `json:"black"`
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

type ChessComPgn struct {
	Event           string
	Site            string
	Date            string
	Round           string
	White           string
	Black           string
	Result          string
	CurrentPosition string
	Timezone        string
	ECO             string
	ECOUrl          string
	UTCDate         string
	UTCTime         string
	WhiteElo        string
	BlackElo        string
	TimeControl     string
	Termination     string
	StartTime       string
	EndDate         string
	EndTime         string
	Link            string
	Moves           string
}
