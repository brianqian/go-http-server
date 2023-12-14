package chess

import (
	"base/internal/http_client"
	"base/types"
	"encoding/json"
	"fmt"
	"log"
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
	resp, _ := cc.client.Get(fmt.Sprintf("https://api.chess.com/pub/player/%s", username))
	profile := &types.ChessComProfile{}
	err := json.Unmarshal(resp, profile)

	if err != nil {
		log.Fatal("")
	}

	cc.profile = profile
	return profile
}
