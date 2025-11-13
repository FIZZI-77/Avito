package service

import "context"

type healthService struct{}

func NewHealthService() Health {
	return &healthService{}
}

func (h *healthService) Ping(ctx context.Context) error {
	return nil
}
