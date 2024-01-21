package api

import (
	"base/internal/chess"
	"base/types"
	"context"
)

type ChessService struct{}

var client = chess.BuildChessClient()

func (c *ChessService) GetProfileByUsername(ctx context.Context, username string) *types.ChessComProfile {
	return client.GetProfileByUsername(username)
}

func (c *ChessService) GetChessComPgn(ctx context.Context, username string, date string) ([]types.RawChessComPgn, error) {
	return client.GetGamesFromChessCom(username, date)
}
