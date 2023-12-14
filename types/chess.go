package types

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
