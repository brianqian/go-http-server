package chess

import (
	"base/internal/http_client"
	"base/types"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"time"
)

type ChessClient struct {
	profile *types.ChessComProfile
	client  types.HttpClient
}

func BuildChessClient() *ChessClient {
	client := &ChessClient{
		profile: &types.ChessComProfile{},
		client:  http_client.BuildHttpClient(),
	}
	return client

}

func (cc *ChessClient) GetProfileByUsername(username string) *types.ChessComProfile {
	resp, _, err := cc.client.Get(fmt.Sprintf("https://api.chess.com/pub/player/%s", username))
	if err != nil {

		slog.Warn("API Error", err)
	}
	var profile = &types.ChessComProfile{}
	err = json.Unmarshal(resp, profile)

	if err != nil {
		slog.Warn("Error unmarshaling profile")
	}

	cc.profile = profile
	return profile
}

func (cc *ChessClient) GetGamesFromChessCom(username string, date string) ([]types.RawChessComPgn, error) {
	// DateOnly   = "2006-01-02"
	d, err := time.Parse(time.DateOnly, date)
	if err != nil {
		fmt.Println("Error parsing, 400!")
		return nil, err
	}
	year, month, _ := d.Date()
	intMonth := int(month)

	var strMonth string
	if intMonth < 10 {
		strMonth = "0" + strconv.Itoa(intMonth)
	} else {
		strMonth = strconv.Itoa(intMonth)
	}
	url := fmt.Sprintf("https://api.chess.com/pub/player/%v/games/%d/%v", username, year, strMonth)

	body, _, err := cc.client.MakeRequest("GET", url, types.RequestOpts{
		Body:    nil,
		Headers: nil,
	})
	if err != nil {
		return nil, err
	}

	var games struct {
		Games []types.RawChessComPgn `json:"games"`
	}

	err = json.Unmarshal(body, &games)
	if err != nil {
		slog.Warn("Error unmarshaling games")
	}

	return games.Games, nil
}
